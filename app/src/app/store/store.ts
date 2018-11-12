// import { combineReducers, Dispatch, Reducer, Action, AnyAction } from 'redux'
import { combineReducers, Reducer } from 'redux'
import { Store, createStore, applyMiddleware } from 'redux'
import thunkMiddleware from 'redux-thunk'

import { BlockchainProps } from './blockchain/types'
import { reducer as blockchainReducer } from './blockchain/reducer'

import { InfoProps } from './info/types'
import { reducer as aboutReducer } from './info/about/reducer'
import { reducer as homeReducer } from './info/home/reducer'
import { reducer as helpReducer } from './info/help/reducer'
import { reducer as overviewReducer } from './info/overview/reducer'

// The top-level state object
export interface ApplicationState {
  blockchain: BlockchainProps
  about: InfoProps
  home: InfoProps
  help: InfoProps
  overview: InfoProps
}

export const rootReducer: Reducer<ApplicationState> = combineReducers<ApplicationState>({
  blockchain: blockchainReducer,
  about: aboutReducer,
  home: homeReducer,
  help: helpReducer,
  overview: overviewReducer
})

export function configureStore(
  initialState: ApplicationState
): Store<ApplicationState> {

  // create the redux-saga middleware
  const store = createStore(
    rootReducer,
    initialState,
    applyMiddleware(thunkMiddleware)
  )

  // Don't forget to run the root saga, and return the store object.
  // sagaMiddleware.run(rootSaga)
  return store
}
