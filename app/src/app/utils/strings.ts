import { Paths as configPaths } from './config'

class App {

  static readonly title='ReportAid'
  static readonly tagline='Using Blockchain Technology to Add Trust to Humanitarian Aid Reporing'
  static readonly copyright='[University of Sussex](https://www.sussex.ac.uk/) © 2018'
  static readonly author='Created by [Steven Huckle](https://glowkeeper.github.io/)'

}

class Paths {

  // AppBar
  static readonly home='Home'
  static readonly blockchain = 'Blockchain Info'
  static readonly about='About'
  static readonly overview='Overview'
  static readonly help='Help'
  static readonly writer='Create Records'
  static readonly reader='Read Records'

  static readonly orgWriter='Organisation'
  static readonly orgReportWriter='Report'
  static readonly orgReportBudgetsWriter = "Report Budget"
  static readonly orgReportExpenditureWriter = "Report Expenditure"
  static readonly orgReportRecipientBudgetsWriter = "Report Recipient Budget"
  static readonly orgReportRegionBudgetsWriter='Report Region Budgets'
  static readonly orgReader='Organisations'
  static readonly orgReportReader='Reports'
  static readonly orgReportBudgetsReader='Report Budgets'
  static readonly orgReportExpenditureReader='Report Expenditure'
  static readonly orgReportRecipientBudgetsReader='Report Recipient Budgets'
  static readonly orgReportRegionBudgetsReader='Report Region Budgets'
}

class Blockchain {

  static heading = 'Blockchain Data'
}

class Home {

  static heading = 'Home'

  static info = `Use this application to record humanitarian aid data and to get humanitarian aid information.<br /><br />Read the [Overview](#${configPaths.overview}) section to learn about the origins of **ReportAid**.<br /><br />The [Help](#${configPaths.help}) section gives brief instructions as to how to use **ReportAid** - in essence, to create a humanitarian aid record, click on the [Create Record](#${configPaths.writer}) link and fill in all fields. To retrieve a humanitarian aid record, click on the [Read Record](#${configPaths.reader}) link. You must have [MetaMask](https://metamask.io/) installed.`
}

class About {

  static heading = 'About ReportAid'

  static info = '**ReportAid** version 0.0.2.<br /><br />Created by [Steven Huckle](https://glowkeeper.github.io/) at the [University of Sussex](https://www.sussex.ac.uk/). '

}

class Overview {

  static heading = 'Overview of ReportAid'

  static info = '**ReportAid** is the result of an academic paper titled: _Humanitarian Aid - a Blockchain Application That Adds Trust to Aid Provisioning_. The article discusses how the trust mechanisms of blockchain technology might be used to promote transparanecy in humanitarian aid provisioning. The general idea is that blockchains can add trust to the reporting of humanitarian aid funding.<br /><br />For more information about **ReportAid**, please contact s dot huckle at sussex dot ac dot uk.'
}

class Help {

  static heading = 'ReportAid Help'

  static info = `**ReportAid** allows humanitarian aid organisations to record information about funding.<br /><br />Have a read of the [Overview](#${configPaths.overview}) section, which gives some background about the app\'. <br /><br />To store a humanitarian aid record, click on the [Create Records](#${configPaths.writer}) link. To retrieve information, click on the [Read Records](#${configPaths.reader}) link. This app' relies on [MetaMask](https://metamask.io/).`
}

class IATIWriter {

  static heading = 'IATI Writer'

  static info = `The [Create Organisation](#${configPaths.orgWriter}) link, lets you create a top-level record for an IATI reporting organsation.`
}

class IATIReader {

  static heading = 'IATI Reader'

  static info = `The [Read Organisations](#${configPaths.orgReader}) link, lets you read a top-level record for an IATI reporting organsation.`
}

class Transaction {

  static heading = "Transaction Details"

  static key = 'Transaction Receipt'
  static summary = 'Transaction Summary'
  static success = 'Successful!'
  static fail = 'Failed!'

}

class Organisation {

  static headingOrgWriter = 'Create Organisation Records'
  static headingOrgReader = 'Read Organisation Records'

  static orgDetails = 'Organisation Details'
  static orgIdentifier = "Organisation Reference"

  static orgName = 'Name'
  static code = 'Code'
  static identifier = 'Identifier'

  static numOrgs = 'Number of Organisations'
}

class OrganisationReport {

  static headingOrgReportWriter = 'Create Organisation Report'
  static headingOrgReportReader = 'Organisation Reports'

  static numOrgReports = "Number of reports"
  static version = "Report Version"
  static orgIdentifier = "Report for Organisation"
  static reportKey = "Report Key"
  static reportingOrgRef = 'Reporting Organisation'
  static reportingOrgType = "Reporting Organisation Type"
  static reportingOrgIsSecondary = "Is Reporting Organisation Secondary?"
  static language = "Report Language"
  static currency = "Report Currency"
  static lastUpdated = "Report Last Updated"


  static numOrganisations = 'Number of Organisations Reporting'
  static orgReportDetails = 'Report Details'
}

class OrganisationReportBudget {

  static headingOrgReportBudgetWriter = 'Create Report Budget'
  static headingOrgReportBudgetReader = 'Report Budgets'

  static numReports = "Number of Organisations Reporting"

  static budgetReference = "Budget Reference"
  static reportReference = "Report Reference"
  static budgetLine = "Budget Line"
  static value = "Value"
  static status = "Status"
  static budgetStart = "Start Date"
  static budgetStartDay = "Start Day"
  static budgetStartMonth = "Start Month"
  static budgetStartYear = "Start Year"
  static budgetEnd = "End Date"
  static budgetEndDay = "End Day"
  static budgetEndMonth = "End Month"
  static budgetEndYear = "End Year"

  static numReportBudgets = "Number of Budgets"
  static reportBudgetDetails = 'Budget Details'
}

class OrganisationReportExpenditure {

  static headingOrgReportExpenditureWriter = 'Create Report Expenditure'
  static headingOrgReportExpenditureReader = 'Report Expenditure'

  static numReports = "Number of Organisations Reporting"

  static expenditureReference = "Expenditure Reference"
  static reportReference = "Report Reference"
  static expenditureLine = "Expenditure Line"
  static value = "Value"
  static status = "Status"
  static expenditureStart = "Start Date"
  static expenditureStartDay = "Start Day"
  static expenditureStartMonth = "Start Month"
  static expenditureStartYear = "Start Year"
  static expenditureEnd = "End Date"
  static expenditureEndDay = "End Day"
  static expenditureEndMonth = "End Month"
  static expenditureEndYear = "End Year"

  static numReportExpenditure = "Number of Expenditure"
  static reportExpenditureDetails = 'Expenditure Details'
}

class OrganisationReportRecipientBudget {

  static headingOrgReportRecipientBudgetWriter = 'Create Report Recipient Budget'
  static headingOrgReportRecipientBudgetReader = 'Report Recipient Budgets'

  static numReports = "Number of Organisations Reporting"

  static budgetReference = "Budget Reference"
  static reportReference = "Report Reference"
  static orgReference = "ReciRecipientpient Reference"
  static budgetLine = "Budget Line"
  static value = "Value"
  static status = "Status"
  static budgetStart = "Start Date"
  static budgetStartDay = "Start Day"
  static budgetStartMonth = "Start Month"
  static budgetStartYear = "Start Year"
  static budgetEnd = "End Date"
  static budgetEndDay = "End Day"
  static budgetEndMonth = "End Month"
  static budgetEndYear = "End Year"

  static numReportBudgets = "Number of Recipient Budgets"
  static reportBudgetDetails = 'Recipient Budget Details'
}

class OrganisationReportRegionBudget {

  static headingOrgReportRegionBudgetWriter = 'Create Report Region Budget'
  static headingOrgReportRegionBudgetReader = 'Report Region Budgets'

  static numReports = "Number of Organisations Reporting"

  static budgetReference = "Budget Reference"
  static reportReference = "Report Reference"
  static regionReference = "Region Reference"
  static budgetLine = "Budget Line"
  static value = "Value"
  static status = "Status"
  static budgetStart = "Start Date"
  static budgetStartDay = "Start Day"
  static budgetStartMonth = "Start Month"
  static budgetStartYear = "Start Year"
  static budgetEnd = "End Date"
  static budgetEndDay = "End Day"
  static budgetEndMonth = "End Month"
  static budgetEndYear = "End Year"

  static numReportBudgets = "Number of Region Budgets"
  static reportBudgetDetails = 'Region Budget Details'

}


export { App,
         Paths,
         Blockchain,
         Home,
         About,
         Overview,
         Help,
         IATIWriter,
         IATIReader,
         Transaction,
         Organisation,
         OrganisationReport,
         OrganisationReportBudget,
         OrganisationReportExpenditure,
         OrganisationReportRecipientBudget,
         OrganisationReportRegionBudget
       }
