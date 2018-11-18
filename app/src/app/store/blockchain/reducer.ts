import { BlockchainActionTypes, BlockchainAction, BlockchainProps } from './types'

const initialState: BlockchainProps = {
  APIName: '',
  networkName: '',
  networkChainId: '',
  networkENSAddress: '',
  account: '',
  web3: {},
  ethers: {},
}

export const blockchainReducer = (state: BlockchainProps = initialState, action: BlockchainAction): BlockchainProps => {
  if ( (action.type == BlockchainActionTypes.ADD_INFO ) ||
       (action.type == BlockchainActionTypes.ADD_ACCOUNT ) ||
       (action.type == BlockchainActionTypes.ADD_OBJECTS )
     ) {
    console.log(action.payload)
    return (<any>Object).assign({}, state, action.payload)
  } else {
    return state
  }
}