// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

interface IERC721 {
    // function ownerOf(uint256 tokenId) external view returns (address owner);
    function transferFrom(address from, address to, uint256 tokenId) external;
}
contract EnglishAuction {
    // events
    event Start(uint startTime, uint endTime);
    event Bid(address indexed bidder, uint value); // indexed to make it easier to filter
    event End(address highestBider, uint highestBid);
    event Withdraw(address bidder, uint value);
    //  acution states
    bool public started;
    bool public ended;
    uint public endedTime;
    uint public highestBid;
    address public highestBider;
    mapping(address => uint) public allBids;

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    // for constructor
    address payable public immutable owner;
    IERC721 public immutable nft;

    uint public immutable nftId;
    // _nftId because one NFT contract can have multiple NFTs and we need to specify which one we are auctioning
    constructor(address _nft, uint _nftId) {
        // init values
        owner = payable(msg.sender);
        // owner and NFT
        // nft follows the IERC721 interface
        nft = IERC721(_nft);
        nftId = _nftId;

    }
    function bid() external payable {
        // validation
        require(started, "Auction not started");
        require(msg.value > highestBid, "Bid too low");
        require(block.timestamp < endedTime, "Auction ended");
        // highest bidder, all bids, highest bidder
        allBids[highestBider] += highestBid;
        highestBid = msg.value;
        highestBider = msg.sender;

        emit Bid(msg.sender, msg.value);
        

    }
    function start(uint _openingBid, uint _duration) external onlyOwner {
        // validation
        require(!started, "Auction already started");

        // update auction state

        highestBid = _openingBid;
        endedTime = block.timestamp + _duration;

        nft.transferFrom(owner, address(this), nftId);
        started = true;

        emit Start(block.timestamp, endedTime);

    }
    function end() external onlyOwner {
        // validations
        require(started, "Auction not started");
        require(block.timestamp >= endedTime, "Auction not ended");
        require(!ended, "Auction already ended");
        // highest bidder receive the item
        // owner receives ether
        ended = true;
        nft.transferFrom(address(this), highestBider, nftId);
        owner.transfer(highestBid);
        emit End(highestBider, highestBid);
    }
    function withdraw() external {
        // bidder to receive fund from the all bids state 
        uint value = allBids[msg.sender];
        allBids[msg.sender] = 0;
        if (value > 0) {
            payable(msg.sender).transfer(value);
        }
        emit Withdraw(msg.sender, value);
    }
}
