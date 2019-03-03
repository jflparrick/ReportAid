const StringsLib = artifacts.require("./Strings.sol");

const IATIBudgets = artifacts.require("./IATIBudgets.sol");

const IATIOrgs = artifacts.require("./IATIOrgs.sol");
const IATIOrganisations = artifacts.require("./IATIOrganisations.sol");
const IATIOrganisation = artifacts.require("./IATIOrganisation.sol");
const IATIOrganisationDocs = artifacts.require("./IATIOrganisationDocs.sol");
const IATIOrganisationBudgets = artifacts.require("./IATIOrganisationBudgets.sol");
const IATIOrganisationExpenditure = artifacts.require("./IATIOrganisationExpenditure.sol");
const IATIOrganisationRecipientBudgets = artifacts.require("./IATIOrganisationRecipientBudgets.sol");
const IATIOrganisationRegionBudgets = artifacts.require("./IATIOrganisationRegionBudgets.sol");
const IATIOrganisationCountryBudgets = artifacts.require("./IATIOrganisationCountryBudgets.sol");

const IATIActivities = artifacts.require("./IATIActivities.sol");
const IATIActivity = artifacts.require("./IATIActivity.sol");
const IATIActivityAdditional = artifacts.require("./IATIActivityAdditional.sol");
const IATIActivityDates = artifacts.require("./IATIActivityDates.sol");
const IATIActivityParticipatingOrgs = artifacts.require("./IATIActivityParticipatingOrgs.sol");

module.exports = function(deployer) {

  let orgsAddress;
  let organisationsAddress;
  let organisationAddress;
  let organisationDocsAddress;
  let organisationBudgetsAddress;
  let organisationExpenditureAddress;
  let organisationRecipientBudgetsAddress;
  let organisationRegionBudgetsAddress;
  let organisationCountryBudgetsAddress;
  let activitiesAddress;
  let activityAddress;
  let activityAdditionalAddress;
  let activityDatesAddress;
  let activityParticipatingOrgsAddress;

  deployer.deploy(StringsLib);
  deployer.link(StringsLib, [IATIBudgets,
                             IATIOrgs,
                             IATIOrganisations,
                             IATIOrganisation,
                             IATIOrganisationDocs,
                             IATIActivities,
                             IATIActivity,
                             IATIActivityAdditional,
                             IATIActivityDates,
                             IATIActivityParticipatingOrgs
                            ]);

  deployer.deploy(IATIOrgs).then(() => {
    orgsAddress = "\"" + IATIOrgs.address + "\"";

  });

  deployer.deploy(IATIOrganisations).then(() => {
    organisationsAddress = "\"" + IATIOrganisations.address + "\"";
  });

  deployer.deploy(IATIOrganisation).then(() => {
    organisationAddress = "\"" + IATIOrganisation.address + "\"";
  });

  deployer.deploy(IATIOrganisationDocs).then(() => {
    organisationDocsAddress = "\"" + IATIOrganisationDocs.address + "\"";
  });

  deployer.deploy(IATIBudgets).then(() => {
    //console.log(IATIBudgets.address)
    deployer.deploy(IATIOrganisationBudgets, IATIBudgets.address).then(() => {
      organisationBudgetsAddress = "\"" + IATIOrganisationBudgets.address + "\"";
    });
    deployer.deploy(IATIOrganisationExpenditure, IATIBudgets.address).then(() => {
      organisationExpenditureAddress = "\"" + IATIOrganisationExpenditure.address + "\"";
    });
    deployer.deploy(IATIOrganisationRecipientBudgets, IATIBudgets.address).then(() => {
      organisationRecipientBudgetsAddress = "\"" + IATIOrganisationRecipientBudgets.address + "\"";
    });
    deployer.deploy(IATIOrganisationRegionBudgets, IATIBudgets.address).then(() => {
      organisationRegionBudgetsAddress = "\"" + IATIOrganisationRegionBudgets.address + "\"";
    });
    deployer.deploy(IATIOrganisationCountryBudgets, IATIBudgets.address).then(() => {
      organisationCountryBudgetsAddress = "\"" + IATIOrganisationCountryBudgets.address + "\"";
    });
  });

  deployer.deploy(IATIActivities).then(() => {
    activitiesAddress = "\"" + IATIActivities.address + "\"";
  });

  deployer.deploy(IATIActivity).then(() => {
    activityAddress = "\"" + IATIActivity.address + "\"";
  });

  deployer.deploy(IATIActivityAdditional).then(() => {
    activityAdditionalAddress = "\"" + IATIActivityAdditional.address + "\"";
  });

  deployer.deploy(IATIActivityDates).then(() => {
    activityDatesAddress = "\"" + IATIActivityDates.address + "\"";
  });

  deployer.deploy(IATIActivityParticipatingOrgs).then(() => {
    activityParticipatingOrgsAddress = "\"" + IATIActivityParticipatingOrgs.address + "\"";
  });

  deployer.then( () => {
    console.log("static orgsAddress =", orgsAddress);
    console.log("static organisationsAddress =", organisationsAddress);
    console.log("static organisationAddress =", organisationAddress);
    console.log("static organisationDocsAddress =", organisationDocsAddress);
    console.log("static organisationBudgetsAddress =", organisationBudgetsAddress);
    console.log("static organisationExpenditureAddress =", organisationExpenditureAddress);
    console.log("static organisationRecipientBudgetsAddress =", organisationRecipientBudgetsAddress);
    console.log("static organisationRegionBudgetsAddress =", organisationRegionBudgetsAddress);
    console.log("static organisationCountryBudgetsAddress =", organisationCountryBudgetsAddress);
    console.log("static activitiesAddress =", activitiesAddress);
    console.log("static activityAddress =", activityAddress);
    console.log("static activityAdditionalAddress =", activityAdditionalAddress);
    console.log("static activityDatesAddress =", activityDatesAddress);
    console.log("static activityParticipatingOrgsAddress =", activityParticipatingOrgsAddress);
  });
};
