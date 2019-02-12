import * as React from 'react'
import { Link } from 'react-router-dom'
import MenuList from '@material-ui/core/MenuList'
import MenuItem from '@material-ui/core/MenuItem'

import IconButton from '@material-ui/core/IconButton'
import Create from '@material-ui/icons/Create'
import List from '@material-ui/icons/List'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../styles/theme'

import { Paths, Organisation } from '../utils/strings'
import { Paths as PathConfig } from '../utils/config'

class Sider extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <div>
        <h3>{Organisation.headingOrganisationWriter}</h3>
        <MenuList>
          <Link to={PathConfig.orgWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgWriter}>
                <Create />
              </IconButton>
              {Paths.orgWriter}
            </MenuItem>
          </Link>
          <Link to={PathConfig.organisationDocsWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.organisationDocsWriter}>
                <Create />
              </IconButton>
              {Paths.organisationDocsWriter}
            </MenuItem>
          </Link>
          <Link to={PathConfig.organisationBudgetsWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.organisationBudgetsWriter}>
                <Create />
              </IconButton>
              {Paths.organisationBudgetsWriter}
            </MenuItem>
          </Link>
          <Link to={PathConfig.organisationExpenditureWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.organisationExpenditureWriter}>
                <Create />
              </IconButton>
              {Paths.organisationExpenditureWriter}
            </MenuItem>
          </Link>
          <Link to={PathConfig.organisationRecipientBudgetsWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.organisationRecipientBudgetsWriter}>
                <Create />
              </IconButton>
              {Paths.organisationRecipientBudgetsWriter}
            </MenuItem>
          </Link>
          <Link to={PathConfig.organisationRegionBudgetsWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.organisationRegionBudgetsWriter}>
                <Create />
              </IconButton>
              {Paths.organisationRegionBudgetsWriter}
            </MenuItem>
          </Link>
          <Link to={PathConfig.organisationCountryBudgetsWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.organisationCountryBudgetsWriter}>
                <Create />
              </IconButton>
              {Paths.organisationCountryBudgetsWriter}
            </MenuItem>
          </Link>
        </MenuList>

        <h3>{Organisation.headingOrganisationReader}</h3>
        <MenuList>
          <Link to={PathConfig.orgsReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgsReader}>
                <List />
              </IconButton>
              {Paths.orgsReader}
            </MenuItem>
          </Link>
          <Link to={PathConfig.organisationReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.organisationReader}>
                <List />
              </IconButton>
              {Paths.organisationReader}
            </MenuItem>
          </Link>
          <Link to={PathConfig.organisationDocsReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.organisationDocsReader}>
                <List />
              </IconButton>
              {Paths.organisationDocsReader}
            </MenuItem>
          </Link>
          <Link to={PathConfig.organisationBudgetsReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.organisationBudgetsReader}>
                <List />
              </IconButton>
              {Paths.organisationBudgetsReader}
            </MenuItem>
          </Link>
          <Link to={PathConfig.organisationExpenditureReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.organisationExpenditureReader}>
                <List />
              </IconButton>
              {Paths.organisationExpenditureReader}
            </MenuItem>
          </Link>
          <Link to={PathConfig.organisationRecipientBudgetsReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.organisationRecipientBudgetsReader}>
                <List />
              </IconButton>
              {Paths.organisationRecipientBudgetsReader}
            </MenuItem>
          </Link>
          <Link to={PathConfig.organisationRegionBudgetsReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.organisationRegionBudgetsReader}>
                <List />
              </IconButton>
              {Paths.organisationRegionBudgetsReader}
            </MenuItem>
          </Link>
          <Link to={PathConfig.organisationCountryBudgetsReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.organisationCountryBudgetsReader}>
                <List />
              </IconButton>
              {Paths.organisationCountryBudgetsReader}
            </MenuItem>
          </Link>
        </MenuList>

      </div>
    )
  }
}

export const SiderOrganisationMenu = withTheme(withStyles(styles)(Sider))
