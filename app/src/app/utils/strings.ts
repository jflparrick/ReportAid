class App {

  static readonly title='ReportAid'
  static readonly tagline='Using Blockchain Technology to Add Trust to Humanitarian Aid Reporing'
  static readonly copyright='University of Sussex © 2018'
  static readonly author='Created by Steven Huckle'

}

class Paths {

  static readonly home='Home'
  static readonly blockchain = 'Blockchain Info\''
  static readonly about='About'
  static readonly overview='Overview'
  static readonly help='Help'
  static readonly writer='Create Record'
  static readonly reader='Fetch Record'
}

class Blockchain {

  static heading = 'Blockchain Data'
}

class Home {

  static heading = 'Home'

  static info = 'Use this application to record humanitarian aid data and to get humanitarian aid information.<br /><br />Read the [About](#/about) section to learn about the origins of **ReportAid**.<br /><br />The [Overview](#/overview) section describes a scenario where an organisation uses **ReportAid** to store information about their funding.<br /><br />The [Help](#/help) section gives brief instructions as to how to use **ReportAid** - in essence, to create a humanitarian aid record, click on the [Create Record](#/create) link and fill in all fields. To retrieve a humanitarian aid record, click on the [Fetch Record](#/fetch) link.'
}

class About {

  static heading = 'About ReportAid'

  static info = '**ReportAid** is the result of an academic paper titled: _Humanitarian Aid - a Blockchain Application That Adds Trust to Aid Provisioning_. The article discusses how the trust mechanisms of blockchain technology might be used to promote transparanecy in humanitarian aid provisioning. The general idea is that blockchains can add trust to the reporting of humanitarian aid funding.<br /><br />For more information about **ReportAid**, please contact s dot huckle at sussex dot ac dot uk.'

}

class Overview {

  static heading = 'Overview of ReportAid'

  static info = 'blah...'
}

class Help {

  static heading = 'ReportAid Help'

  static info = '**ReportAid** allows humanitarian aid organisations to record information about funding.<br /><br />Have a read of the [About](#/about) section, which gives some background to the app\'. The [Overview](#/overview) section describes a scenario where an organisation uses **ReportAid** to store information about their funding.<br /><br />To store a humanitarian aid record, click on the [Create Record](#/create) link and fill in all fields. That will create some blockchain transactions, which can be signed in [MetaMask](https://metamask.io/). To retrieve information, click on the [Fetch Record](#/fetch) link.'
}

class IATIWriter {

  static heading = 'IATI Writer'

  static info = 'Here, you can create IATI information.<br /><br />The [create organisation](#/create-organisation) link, lets you create a top-level record for an IATI reporting organsation.'
}

class IATIReader {

  static heading = 'IATI Reader'

  static info = 'Here, you can fetch IATI information.<br /><br />The [fetch organisation](#/fetch-organisation) link, lets you read a top-level record for an IATI reporting organsation.'
}

class Organisation {

  static orgName = 'Name'
  static identifier = 'Identifier'
  static type = 'Type'
}

export { App, Paths, Blockchain, Home, About, Overview, Help, IATIWriter, IATIReader, Organisation }
