pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

enum IATIElement {
    NONE,
    ORGS,
    ACTIVITIES,
    ACTIVITY,
    ACTIVITYADDITIONAL,
    ACTIVITYBUDGETS,
    ACTIVITYDATES,
    ACTIVITYPARTICIPATINGORGS,
    ACTIVITYRELATEDACTIVITIES,
    ACTIVITYSECTORS,
    ACTIVITYTERRITORIES,
    ACTIVITYTRANSACTIONS,
    ORGANISATIONS,
    ORGANISATION,
    ORGANISATIONBUDGETS,
    ORGANISATIONCOUNTRYBUDGETS,
    ORGANISATIONDOCS,
    ORGANISATIONEXPENDITURE,
    ORGANISATIONRECIPIENTBUDGETS,
    ORGANISATIONREGIONBUDGETS,
    MAX
}

enum Scope {
    NONE,
    GLOBAL,
    REGIONAL,
    MULTINATIONAL,
    NATIONAL,
    SUBNATIONALMULTIFIRSTLEVEL,
    SUBNATIONALSINGLEFIRSTLEVEL,
    SUBNATIONALSINGLESECONDLEVEL,
    SINGLELOCATION,
    MAX
}

enum CollaborationType {
    NONE,
    BILATERAL,
    MULTILATERALINFLOWS,
    BILATERALPRIVATE,
    MULTILATERALOUTFLOWS,
    PRIVATEOUTFLOWS,
    BILATERALCOREFUNDED,
    TRIANGULAR,
    MAX
}

enum TiedStatus {
    UNSPECIFIED,
    UNSPECIFIEDONE,
    NONE,
    PARTIALLYTIED,
    TIED,
    UNTIED,
    MAX
}

enum DateCodes {
    NONE,
    PLANNEDSTART,
    ACTUALSTART,
    PLANNEDEND,
    ACTUALEND,
    MAX
}

enum Role {
    NONE,
    FUNDING,
    ACCOUNTABLE,
    EXTENDING,
    IMPLEMENTING,
    MAX
}

enum RelationType {
    NONE,
    PARENT,
    CHILD,
    SIBLING,
    COFUNDED,
    THIRDPARTY,
    MAX
}

enum ActivityHierarchy {
    NONE,
    ACTIVITY,
    SUBACTIVITY,
    SUBSUBACTIVITY,
    MAX
}

enum ActivityStatus {
    NONE,
    PIPELINEIDENTIFICATION,
    IMPLEMENTATION,
    FINALISATION,
    CLOSED,
    CANCELLED,
    SUSPENDED,
    MAX
}

enum BudgetNotProvided {
    NONE,
    COMMERCIALRESTRICTIONS,
    LEGALRESTRICTIONS,
    RAPIDONSETEMERGENCY,
    MAX
}

enum TransactionType {
    NONE,
    INCOMINGFUNDS,
    OUTGOINGCOMMITMENT,
    DISBURSEMENT,
    EXPENDITURE,
    INTERESTPAYMENT,
    LOANREPAYMENT,
    REIMBURSEMENT,
    PURCHASEOFEQUITY,
    SALEOFEQUITY,
    CREDITGUARANTEE,
    INCOMINGCOMMITMENT,
    OUTGOINGPLEDGE,
    INCOMINGPLEDGE,
    MAX
}

enum DisbursementChannel {
    NONE,
    TREASURY,
    BANKACCOUNT,
    NGO,
    DONORMANAGED,
    MAX
}

enum BudgetOwner {
    NONE,
    ORGEXPENDITURE,
    ORG,
    ORGRECIPIENT,
    ORGREGION,
    ORGCOUNTRY,
    ACTIVITY,
    MAX
}

enum BudgetStatus {
    NONE,
    INDICATIVE,
    COMMITTED,
    MAX
}

enum BudgetType {
    NONE,
    ORIGINAL,
    REVISED,
    MAX
}

enum DocAttributes {
    TITLE,
    FORMAT,
    URL,
    CATEGORY,
    COUNTRYCODE,
    DESC,
    LANG,
    DATE
}