import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getOrgs } from '../../../store/IATI/IATIReader/organisations/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgData } from '../../../store/IATI/IATIReader/organisations/types'

import { Organisation as OrgStrings } from '../../../utils/strings'

import { getDictEntries } from '../../../components/io/dict'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgProps {
  num: number
  orgs: OrgData
}

interface OrgDispatchProps {
  getOrgs: () => void
}

type OrgReaderProps =  WithStyles<typeof styles> & OrgProps & OrgDispatchProps

export class Orgs extends React.Component<OrgReaderProps> {

  constructor (props: OrgReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getOrgs()
  }

  render() {

    const orgsData = Object.keys(this.props.orgs)
    let xs = ""
    if ( orgsData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      orgsData.forEach((key) => {
        xs += `**${OrgStrings.orgIdentifier}**: ${key}<br />`
        const values = Object.values(this.props.orgs[key])
        //console.log('Values: ', values)
        xs += `**${OrgStrings.numOrgs}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((thisKey) => {
          //console.log(': ', values[1][thisKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][thisKey].hasOwnProperty('version') && values[1][thisKey].version != "" ) {
            const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
            const language =  ethers.utils.parseBytes32String(values[1][thisKey].lang)
            const currency =  ethers.utils.parseBytes32String(values[1][thisKey].currency)
            const lastUpdated =  ethers.utils.parseBytes32String(values[1][thisKey].lastUpdatedTime)
            xs+= `**${OrgStrings.reportingOrgRef}**: ${values[1][thisKey].reportingOrg.orgRef} <br />`
            xs+= `**${OrgStrings.reportKey}**: ${thisKey} <br />`
            xs+= `**${OrgStrings.reportingOrgType}**: ${values[1][thisKey].reportingOrg.orgType} <br />`
            xs+= `**${OrgStrings.reportingOrgIsSecondary}**: ${values[1][thisKey].reportingOrg.isSecondary} <br />`
            xs+= `**${OrgStrings.version}**: ${version} <br />`
            xs+= `**${OrgStrings.language}**: ${language} <br />`
            xs+= `**${OrgStrings.currency}**: ${currency} <br />`
            xs+= `**${OrgStrings.lastUpdated}**: ${lastUpdated} <br /><br />`
          }
        })
        length += 1
        length == orgsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrgStrings.headingOrgReader}</h2>
        <p>
          <b>{OrgStrings.numOrganisations}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgStrings.orgDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgsReader.num,
    orgs: state.orgsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    getOrgs: () => dispatch(getOrgs())
  }
}

export const Organisations = withTheme(withStyles(styles)(connect<OrgProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Orgs)))