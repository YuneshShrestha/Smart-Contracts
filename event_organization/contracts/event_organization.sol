// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract EventContract {
    struct Event{
        address organizer;
        string name;
        uint date;
        uint price;
        uint ticketCount;
        uint ticketRemain;
    }
    mapping(uint=>Event) public events;
    mapping(address=>mapping (uint=>uint)) public tickets;
    uint public nextId;

    function createEvent(string memory name, uint date, uint price, uint ticketCount) external {
        require(date > block.timestamp, "date is before now");
        require(ticketCount>0, "ticket count is not correct");
        events[nextId] = Event(msg.sender, name,  date,  price,  ticketCount,  ticketCount);
        nextId++;
    }
    modifier isTicketAvailable(uint quantity, uint eventId){
         require(quantity>0, "ticket count is not correct");
        require(events[eventId].date != 0, "event doesn't exist");
        require(events[eventId].ticketCount > quantity, "ticket count is not correct");
        _;
    }
    function buyTicket(uint eventId, uint quantity) isTicketAvailable(quantity, eventId) external payable{
       
        require(events[eventId].date >= block.timestamp,"Event time has exceeded");

        Event storage _event = events[eventId];

        require(msg.value>= (_event.price*quantity),"not enough amount");
        _event.ticketRemain -= quantity;
        tickets[msg.sender][eventId] += quantity;

    }
    function transferTicket(uint eventId, uint quantity, address to) isTicketAvailable(quantity, eventId) external{
        tickets[msg.sender][eventId] -= quantity;
        tickets[to][eventId] += quantity;
    }
}