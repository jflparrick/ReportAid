import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../../types'
import { OrganisationRegionBudgetProps, IATIOrganisationRegionBudgetProps } from '../../../types'
import { IATIWriterActionTypes } from '../types'

export const setRegionBudget = (details: OrganisationRegionBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const regionBudgetsContract = state.chainContracts.data.contracts.organisationRegionBudgets

    const start = new Date(details.startYear + '/' + details.startMonth + '/' + details.startDay)
    const end = new Date(details.endYear + '/' + details.endMonth + '/' + details.endDay)

    let budgetRef = details.budgetRef
    if ( budgetRef == "" ) {
      budgetRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const regionBudget: IATIOrganisationRegionBudgetProps = {
      regionRef: details.regionRef,
      budgetLine: ethers.utils.formatBytes32String(details.budgetLine),
      finance: {
        value: details.value,
        status: details.status,
        start: ethers.utils.formatBytes32String(start.toISOString()),
        end: ethers.utils.formatBytes32String(end.toISOString())
      }
    }

    //console.log('RegionBudget: ', regionBudget, ' Contract ', orgRegionBudgetsContract)
    let actionType = IATIWriterActionTypes.RECIPIENTREGIONBUDGET_FAILURE
    let txData: TxData = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await regionBudgetsContract.setRegionBudget(details.organisationsRef,
                                                             details.organisationRef,
                                                             budgetRef,
                                                             regionBudget)
      const key = tx.hash
      txData[key] = tx
      actionType = IATIWriterActionTypes.RECIPIENTREGIONBUDGET_SUCCESS
    } catch (error) {
      txData[-1] = txData
      console.log('setRegionBudget error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}