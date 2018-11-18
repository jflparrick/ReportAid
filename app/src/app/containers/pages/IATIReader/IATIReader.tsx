import * as React from 'react'
import { connect } from 'react-redux'
// import { Dispatch } from 'redux'
import { MarkdownText } from '../../../components/io/markdownText'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

import { ApplicationState } from '../../../store'
// import { fetchRequest } from '../../../store/IATIReader/actions'


// We can use `typeof` here to map our dispatch types to the props, like so.
/* interface PropsFromDispatch {
  fetchRequest: typeof fetchRequest
} */

export interface IATIReaderProps {
  title: string
  data: string
}

// Combine both state + dispatch props - as well as any props we want to pass - in a union type.
// type AllProps = IATIReaderProps // & PropsFromDispatch

class IATIReader extends React.Component<WithStyles<typeof styles> & IATIReaderProps> {

  render() {
    return (
      <div>
        <h2>{this.props.title}</h2>
        <MarkdownText text={this.props.data} />
      </div>
    )
  }
}

// It's usually good practice to only include one context at a time in a connected component.
// Although if necessary, you can always include multiple contexts. Just make sure to
// separate them from each other to prevent prop conflicts.
const mapStateToProps = (state: ApplicationState, ownProps: IATIReaderProps): IATIReaderProps => {
  return {
    title: ownProps.title,
    data: ownProps.data
  }
}

// mapDispatchToProps is especially useful for constraining our actions to the connected component.
// You can access these via `this.props`.
/* const mapDispatchToProps = (dispatch: Dispatch) => ({
  fetchRequest: () => dispatch(fetchRequest())
})*/

// Now let's connect our component!
// With redux v4's improved typings, we can finally omit generics here.
export default withTheme(withStyles(styles)(connect(
  mapStateToProps,
  //mapDispatchToProps
)(IATIReader)))
