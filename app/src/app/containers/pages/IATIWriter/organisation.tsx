import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import FormControl from '@material-ui/core/FormControl'
//import { Select } from 'formik-material-ui'
import { TextField, Select } from "material-ui-formik-components"

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrganisationProps } from '../../../store/IATI/types'

import { initialise } from '../../../store/IATI/IATIWriter/organisations/actions'
import { setOrganisation } from '../../../store/IATI/IATIWriter/organisations/organisation/actions'

import { OrganisationsPicker } from '../../../components/io/organisationsPicker'
import { OrganisationPicker } from '../../../components/io/organisationPicker'
import { OrgPicker } from '../../../components/io/orgPicker'
import { TransactionHelper } from '../../io/transactionHelper'

import { Organisation as OrganisationStrings } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const organisationSchema = Yup.object().shape({
  organisationsRef: Yup
    .string()
    .required('Required'),
  organisationRef: Yup
    .string()
    .required('Required'),
  orgRef: Yup
    .string()
    .required('Required'),
  reportingOrgRef: Yup
    .string()
    .required('Required'),
  reportingOrgType: Yup
    .number()
    .required('Required'),
  reportingOrgIsSecondary: Yup
    .boolean()
    .required('Required'),
  lang: Yup
    .string()
    .required('Required'),
  currency: Yup
    .string()
    .required('Required'),
})

interface OrganisationKeyProps {
  organisationsRef: string
  organisationRef: string
}

export interface OrganisationDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
}

type OrganisationFormProps = WithStyles<typeof styles> & OrganisationKeyProps & OrganisationDispatchProps

export class OrganisationForm extends React.Component<OrganisationFormProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrganisationFormProps) {
   super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  handleSubmit = (values: OrganisationProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
    this.props.initialise()
  }

  render() {

    return (
      <div>
        <h2>{OrganisationStrings.headingOrganisationWriter}</h2>
        <div>
          <Formik
            initialValues={ {organisationsRef: "",
                             organisationRef: this.props.organisationRef,
                             orgRef: "",
                             reportingOrgRef: "",
                             reportingOrgType: 0,
                             reportingOrgIsSecondary: false,
                             lang: "",
                             currency: ""
                            }}
            enableReinitialize={true}
            validationSchema={organisationSchema}
            onSubmit={(values: OrganisationProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrganisationProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <Field
                    readOnly
                    name='organisationRef'
                    value={this.props.organisationRef}
                    label={OrganisationStrings.organisationReference}
                    component={TextField}
                  />
                  <OrganisationsPicker
                    setValue={formProps.setFieldValue}
                    name='organisationsRef'
                    label={OrganisationStrings.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrgPicker
                    name='orgRef'
                    label={OrganisationStrings.orgRef}
                  />
                  <ErrorMessage name='orgRef' />
                  <OrgPicker
                    name='reportingOrgRef'
                    label={OrganisationStrings.reportingOrgRef}
                  />
                  <ErrorMessage name='reportingOrgRef' />
                  <Field
                    name="reportingOrgType"
                    label={OrganisationStrings.reportingOrgType}
                    component={Select}
                    options={Helpers.organisationCodes}
                  />
                  <ErrorMessage name='reportingOrgType' />
                  <Field
                    name="reportingOrgIsSecondary"
                    label={OrganisationStrings.reportingOrgIsSecondary}
                    component={Select}
                    options={Helpers.isSecondary}
                  />
                  <ErrorMessage name='reportingOrgIsSecondary' />
                  <Field
                    name="lang"
                    label={OrganisationStrings.language}
                    component={Select}
                    options={Helpers.languageCodes}
                  />
                  <ErrorMessage name='lang' />
                  <Field
                    name="currency"
                    label={OrganisationStrings.currency}
                    component={Select}
                    options={Helpers.currencyCodes}
                  />
                  <ErrorMessage name='currency' />
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
        <TransactionHelper
          submitFunc={this.state.submitFunc}
          resetFunc={this.state.resetFunc}
        />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationKeyProps => {
  //console.log(state.orgReader)
  return {
    organisationsRef: state.keys.data.organisations,
    organisationRef: state.keys.data.newKey
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisation(ownProps)),
    initialise: () => dispatch(initialise())
  }
}

export const Organisation = withTheme(withStyles(styles)(connect<OrganisationKeyProps, OrganisationDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrganisationForm)))
