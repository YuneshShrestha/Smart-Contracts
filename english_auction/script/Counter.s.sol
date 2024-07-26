// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {EnglishAuction} from "../src/English-Auction.sol";

contract EnglishAuctionScript is Script {
    EnglishAuction public englishAuction;

    function setUp() public {}

    function run() public {
        vm.startBroadcast();

        englishAuction = new EnglishAuction(address(this), 0);

        vm.stopBroadcast();
    }
}
