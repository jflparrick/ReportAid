import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'

import { ApplicationState } from '../../store'
import { OrgAddAction, OrganisationProps } from '../../store/IATIWriter/organisationWriter/types'
import { setOrganisation } from '../../store/IATIWriter/organisationWriter/actions'

import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import { TextField } from 'formik-material-ui'

import { Organisation } from '../../utils/strings'


export interface OrgDispatchProps {
  handleSubmit: (values: any) => void
}

const organisationSchema = Yup.object().shape({
  name: Yup
    .string()
    .required('Required'),
  reference: Yup
    .string()
    .required('Required'),
  type: Yup
    .string()
    .required('Required'),
})

type OrgWriterFormProps = OrgDispatchProps

export const OrgForm: React.SFC<OrgWriterFormProps> = (props: OrgWriterFormProps) => {
  return (
    <div>
      <Formik
        initialValues={ {name: '', reference: '', type: ''} }
        validationSchema={organisationSchema}
        onSubmit={(values: OrganisationProps, actions: any) => {
          props.handleSubmit(values)
          actions.setSubmitting(false)
          actions.resetForm()
        }}
        render={({isSubmitting}: FormikProps<any>) => (
          <Form>
            <Field
              name='name'
              label={Organisation.orgName}
              component={TextField}
            />
            <ErrorMessage
              name='name'
            />
            <br />
            <Field
              name='reference'
              label={Organisation.reference}
              component={TextField}
            />
            <ErrorMessage
              name='reference'
            />
            <br />
            <Field
              name='type'
              label={Organisation.type}
              component={TextField}
            />
            <ErrorMessage
              name='type'
            />
            <br />
            {isSubmitting && <LinearProgress />}
            <br />
            <Button
              type='submit'
              variant="raised"
              color="primary"
              disabled={isSubmitting}
            >
              Submit
            </Button>
          </Form>
        )}
      />
    </div>
  )
}

/*const mapStateToProps = (state: ApplicationState): OrganisationProps => {
  return {
    name: state.orgForm.name,
    reference: state.orgForm.reference,
    type: state.orgForm.type
  }
}*/

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, OrgAddAction>): OrgDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisation(ownProps))
  }
}

export const OrgWriterForm = connect<{}, OrgDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrgForm)
