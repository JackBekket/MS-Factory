
pragma solidity ^0.5.11;

import './zeppeline/ownership/Ownable.sol';

contract Deposit is Ownable {

/*
Events
*/

// this event invoke when someone want to cashout funds to plastic card
// this event should been catched by 'cashier' service and procced tx out at Fiat payment processer
// amount is in Wei, so cashier probably need to convert it
event cashOutRequestEvent(string indexed destination, address indexed user, uint amount);

// TODO add indexed payload to event

event cashInRequestEvent(address indexed user, uint amount, string indexed uuid);



/*
Constants
*/


// fiat uuid = request
mapping (string => IRequest) public InRequest;

//mapping (string => bool) public Executed

struct IRequest {

    string paymentType;
    string fiat_uuid;       // ID of transaction at unitpay side
    uint amount;
    address payable user_wallet; // In unitpay system it's params[account]
   // string fiat_address; -- we are not processing this info

    address submited_by; // moonshard operator
    bool executed;
    string payload; // should not be secret info
}

constructor() public payable {

}




// cash out
function cashOutRequest(string memory destination, address user) public payable {
    uint amount = msg.value;
    // FIXME: add amount conversion
    emit cashOutRequestEvent(destination, user, amount);

}

// cash in
// cashier submit request for cash in while getting events from Fiat payment processor
function cashInRequest(address payable user, string memory uuid, uint amount) public onlyOwner {

    emit cashInRequestEvent(user,amount,uuid);

    IRequest memory irq;
    irq.fiat_uuid = uuid;
    irq.user_wallet = user;
   // irq.fiat_address = from;
    irq.submited_by = msg.sender;
    irq.amount = amount;
    irq.executed = false;

    InRequest[uuid] = irq;


}



// TODO Validator key can be added here to prove (submit) transaction from FIAT processor. 
// It 's not neccerily, as we could just use the blockchain validation itself, but do it in more transcendent way
// FIXME : change onlyOwner to onlyValidator
function cashInSubmit(string memory uuid) public onlyOwner {
    IRequest memory irq;
    irq = InRequest[uuid];
    require(irq.executed = false, "transaction is already executed! (reentrancy guard)");
    address payable _user = irq.user_wallet;
    uint amount = irq.amount;
    // Do some conversion for amount (FIXME)
    // ...

    // FIXME : use internal proceedTransaction instead
    _user.transfer(amount);
    irq.executed = true;
    InRequest[uuid] = irq;


}

function proceedTransaction(IRequest memory ts) internal {
    address payable _user = ts.user_wallet;
    uint amount = ts.amount;
    _user.transfer(amount);
}





// fallback
function() external payable {
  //  cashOutRequest()
  // don't allow user to cashOut without pointing destination
  revert("can't cashOut without out address");
}


}