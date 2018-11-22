import { action } from 'typesafe-actions'
import { IATIReaderActionTypes } from './types'

// Here we use the `action` helper function provided by `typesafe-actions`.
// This library provides really useful helpers for writing Redux actions in a type-safe manner.
// For more info: https://github.com/piotrwitek/typesafe-actions
export const fetchRequest = () => action(IATIReaderActionTypes.FETCH_REQUEST)

// Remember, you can also pass parameters into an action creator. Make sure to
// type them properly as well.
export const fetchSuccess = (data: []) => action(IATIReaderActionTypes.FETCH_SUCCESS, data)
export const fetchError = (message: string) => action(IATIReaderActionTypes.FETCH_ERROR, message)
