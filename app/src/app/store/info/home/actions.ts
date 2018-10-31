import { Action, ActionCreator } from 'redux'
import { HomeActionTypes, InfoProps } from '../types'

export interface RequestDataAction extends Action {
  type: HomeActionTypes.REQ_DATA
  payload: InfoProps
}

export const fetchRequest: ActionCreator<RequestDataAction> = (type, payload) => {
  console.log('blah blah blah', type, payload)
  return (
    { type: type, payload: payload }
  )
}