export interface IATIOrgProps {
  name: string
  identifier: string
}

export interface OrganisationProps {
  name: string
  code: string
  identifier: string
}

export interface OrgReportProps {
  orgIdentifier: string
  reportingOrgIdentifier: string
  version: string
}
