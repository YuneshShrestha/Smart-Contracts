const EventOrganization = artifacts.require("EventContract");

module.exports = function(deployer) {
    deployer.deploy(EventOrganization);
}
