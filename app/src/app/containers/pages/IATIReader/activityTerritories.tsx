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

import { getDictEntries } from '../../../components/io/dict'
import { ActivitiesPicker } from '../../../components/io/activitiesPicker'
import { ActivityPicker } from '../../../components/io/activityPicker'
import { FormData } from '../../../store/helpers/forms/types'

import { initialise } from '../../../store/IATI/IATIReader/actions'
import { setFormFunctions } from '../../../store/helpers/forms/actions'
import { getActivityTerritories } from '../../../store/IATI/IATIReader/activities/activityTerritories/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIActivityTerritoryReport, ActivitiesReportProps } from '../../../store/IATI/types'

import { ActivityTerritories as ActivityTerritoriesStrings } from '../../../utils/strings'

const reportSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  activityRef: Yup
    .string()
    .required('Required')
})

interface ActivityTerritoriesProps {
  submittingFunc: Function,
  resettingFunc: Function
  activitiesRef: string,
  activityRef: string,
  territories: IATIActivityTerritoryReport
}

interface ActivityTerritoriesDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivityTerritoriesReaderProps = ActivityTerritoriesProps & ActivityTerritoriesDispatchProps

class Territories extends React.Component<ActivityTerritoriesReaderProps> {

  constructor (props: ActivityTerritoriesReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  componentDidUpdate(previousProps: ActivityTerritoriesReaderProps) {
    if(previousProps.territories != this.props.territories) {
    this.props.submittingFunc(false)
    this.props.resettingFunc()
    }
  }

  handleSubmit = (values: ActivitiesReportProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.initialise()
    this.props.handleSubmit(values)
  }

  render() {

    const xs = getDictEntries(this.props.territories)

    return (
      <div>
        <h2>{ActivityTerritoriesStrings.headingActivityTerritoriesReader}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: "",
                             activityRef: ""
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: ActivitiesReportProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivitiesReportProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <ActivitiesPicker
                    setValue={formProps.setFieldValue}
                    name='activitiesRef'
                    label={ActivityTerritoriesStrings.activitiesReference}
                  />
                  <ErrorMessage name='activitiesRef' />
                  <ActivityPicker
                    setValue={formProps.setFieldValue}
                    name='activityRef'
                    label={ActivityTerritoriesStrings.activityReference}
                  />
                  <ErrorMessage name='activityRef' />
                  <br />
                  {formProps.isSubmitting && <LinearProgress />}
                  <br />
                  <Button type='submit' variant="contained" color="primary" disabled={formProps.isSubmitting}>
                    Submit
                  </Button>
                </FormControl>
              </Form>
            )}
          />
        </div>
        <hr />
        <h3>{ActivityTerritoriesStrings.territoriesDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): ActivityTerritoriesProps => {
  //console.log(state.orgReader)
  return {
    submittingFunc: state.forms.data.submitFunc,
    resettingFunc: state.forms.data.resetFunc,
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity,
    territories: state.report.data as IATIActivityTerritoryReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityTerritoriesDispatchProps => {
  return {
    handleSubmit: (values: ActivitiesReportProps) => dispatch(getActivityTerritories(values)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const ActivityTerritories = connect<ActivityTerritoriesProps, ActivityTerritoriesDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Territories)
