import { OrgWriterActionTypes } from './types'
import { ActionProps, TxProps, TxData } from '../../../types'

const initialState: TxProps = {
  data: {
    '': {}
  }
}

export const reducer = (state: TxProps = initialState, action: ActionProps): TxProps => {
  if ( (action.type == OrgWriterActionTypes.ADD_SUCCESS ) ||
       (action.type == OrgWriterActionTypes.ADD_FAILURE )
     ) {
    //console.log('Orgstate: ', state, action)
    //const data: OrgWriterProps = {...state, ...action.payload.data}
    //console.log('TxData: ', data)
    const data = (action.payload.data as TxData)
    return {...state, ...data}
  } else {
    return state
  }
}
