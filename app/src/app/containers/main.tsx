import * as React from 'react'
import logo from '../images/owl.png'
import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles';
import { withTheme, styles } from '../styles/theme'

import Content from './content'
import AppMenu from './appMenu'

import { AppStrings } from '../utils/strings'

class Main extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <Paper className={this.props.classes.root}>
        <Paper className={this.props.classes.header}>
          <Grid container justify='center' spacing={0}>
            <Grid item xs={12} sm={2}>
               <Paper className={this.props.classes.sider}><img src={logo}/></Paper>
            </Grid>
            <Grid item xs={12} sm={10}>
              <Paper className={this.props.classes.header}><h1>{AppStrings.appTitle}</h1></Paper>
            </Grid>
          </Grid>
        </Paper>
        <Paper className={this.props.classes.content}>
          <Grid container spacing={0}>
            <Grid item xs={12} sm={2}>
              <Paper className={this.props.classes.sider}>
                <AppMenu />
              </Paper>
            </Grid>
            <Grid item xs={12} sm={10}>
              <Paper className={this.props.classes.content}>
                <Content />
              </Paper>
            </Grid>
          </Grid>
        </Paper>
        <Paper className={this.props.classes.footer}>
          <Grid container spacing={0}>
            <Grid item xs={2}>
              <Paper className={this.props.classes.footer}>
                <p>{AppStrings.author}</p>
              </Paper>
            </Grid>
            <Grid item xs={10}>
              <Paper className={this.props.classes.footer}>
                <p>{AppStrings.copyright}</p>
              </Paper>
            </Grid>
          </Grid>
        </Paper>
      </Paper>
    )
  }
}

export default withTheme(withStyles(styles)(Main))
