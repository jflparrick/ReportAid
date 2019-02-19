import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
//import { Date } from 'formik-material-ui'
import FormControl from '@material-ui/core/FormControl'

import { OrganisationsPicker } from '../../../components/io/organisationsPicker'
import { OrganisationPicker } from '../../../components/io/organisationPicker'
import { FormData } from '../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../store/helpers/forms/actions'
import { initialise, getRecipientBudgets } from '../../../store/IATI/IATIReader/organisations/organisationRecipientBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrganisationRecipientBudgetReport } from '../../../store/IATI/IATIReader/organisations/organisationRecipientBudgets/types'
import { OrganisationsReportProps } from '../../../store/IATI/IATIReader/organisations/types'

import { OrganisationRecipientBudget as OrganisationRecipientBudgetStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  organisationsRef: Yup
    .string()
    .required('Required'),
  organisationRef: Yup
    .string()
    .required('Required')
})

interface OrganisationRecipientBudgetProps {
  submittingFunc: Function,
  resettingFunc: Function
  organisationsRef: string,
  organisationRef: string,
  budgets: IATIOrganisationRecipientBudgetReport
}

interface OrganisationRecipientBudgetDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
}

type OrganisationRecipientBudgetsReaderProps =  WithStyles<typeof styles> & OrganisationRecipientBudgetProps & OrganisationRecipientBudgetDispatchProps

class RecipientBudgets extends React.Component<OrganisationRecipientBudgetsReaderProps> {

  constructor (props: OrganisationRecipientBudgetsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  componentDidUpdate(previousProps: OrganisationRecipientBudgetsReaderProps) {
    if(previousProps.budgets != this.props.budgets) {
      this.props.submittingFunc(false)
      this.props.resettingFunc()
    }
  }

  handleSubmit = (values: OrganisationsReportProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.initialise()
    this.props.handleSubmit(values)
  }

  handleOrganisationsChange = (value: string) => {
    this.setState({organisationsRef: value})
  }

  handleOrganisationChange = (value: string) => {
    console.log(value)
  }

  render() {

    const budgetsData = this.props.budgets
    let xs = ""
    let num = 0
    Object.keys(budgetsData).forEach((organisationsKey) => {
      //numOrganisations += 1
      xs += `**${OrganisationRecipientBudgetStrings.organisationsReference}**: ${organisationsKey}<br />`
      Object.keys(budgetsData[organisationsKey].data).forEach((organisationKey) => {

        xs += `**${OrganisationRecipientBudgetStrings.organisationReference}**: ${organisationKey}<br />`
        Object.keys(budgetsData[organisationsKey].data[organisationKey].data).forEach((budgetKey) => {
          if ( budgetsData[organisationsKey].data[organisationKey].data[budgetKey].hasOwnProperty('budgetLine') &&
               budgetsData[organisationsKey].data[organisationKey].data[budgetKey].budgetLine != "" ) {

            num += 1
            const thisbudgetData =  budgetsData[organisationsKey].data[organisationKey].data[budgetKey]

            const budgetLine = ethers.utils.parseBytes32String(thisbudgetData.budgetLine)
            const start = ethers.utils.parseBytes32String(thisbudgetData.finance.start)
            const end = ethers.utils.parseBytes32String(thisbudgetData.finance.end)
            xs+= `**${OrganisationRecipientBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.orgReference}**: ${thisbudgetData.recipientOrgRef} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.value}**: ${thisbudgetData.finance.value} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.status}**: ${thisbudgetData.finance.status} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
      })
    })

    return (
      <div>
        <h2>{OrganisationRecipientBudgetStrings.headingOrganisationRecipientBudgetReader}</h2>
        <div>
          <Formik
            initialValues={ {organisationsRef: "",
                             organisationRef: ""
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: OrganisationsReportProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrganisationsReportProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrganisationsPicker
                    setValue={formProps.setFieldValue}
                    name='organisationsRef'
                    label={OrganisationRecipientBudgetStrings.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    setValue={formProps.setFieldValue}
                    name='organisationRef'
                    label={OrganisationRecipientBudgetStrings.organisationReference}
                  />
                  <ErrorMessage name='organisationRef' />
                  <br />
                  {formProps.isSubmitting && <LinearProgress />}
                  <br />
                  <Button type='submit' variant="raised" color="primary" disabled={formProps.isSubmitting}>
                    Submit
                  </Button>
                </FormControl>
              </Form>
            )}
          />
        </div>
        <hr />
        <p>
          <b>{OrganisationRecipientBudgetStrings.numBudgets}</b>: {num}
        </p>
        <h3>{OrganisationRecipientBudgetStrings.organisationRecipientBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationRecipientBudgetProps => {
  //console.log(state.orgReader)
  return {
    submittingFunc: state.forms.data.submitFunc,
    resettingFunc: state.forms.data.resetFunc,
    organisationsRef: state.keys.data.organisations,
    organisationRef: state.keys.data.organisation,
    budgets: state.organisationRecipientBudgetsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationRecipientBudgetDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(getRecipientBudgets(ownProps)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const OrganisationRecipientBudgets = withTheme(withStyles(styles)(connect<OrganisationRecipientBudgetProps, OrganisationRecipientBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(RecipientBudgets)))
