pragma solidity ^0.5.11;

import '../../zeppeline/token/ERC721/ERC721Enumerable.sol';
import '../../zeppeline/token/ERC721/ERC721Mintable.sol';
import "../../zeppeline/drafts/Counters.sol";







// Ticket is ERC721 (NFT) token with  availability to reedem tickets

/**

    ERC-721 is non-fungible token.  So, each tokenID is a unique ticket

    Usefull tips:
    // Mapping from owner to list of owned token IDs
    mapping(address => uint256[]) private _ownedTokens;

    // Mapping from token ID to index of the owner tokens list
    mapping(uint256 => uint256) private _ownedTokensIndex;

    // Array with all token ids, used for enumeration
    uint256[] private _allTokens;

    // Mapping from token id to position in the allTokens array
    mapping(uint256 => uint256) private _allTokensIndex;

    Gets the list of token IDs of the requested owner.
     function _tokensOfOwner(address owner) internal view returns (uint256[] storage) {
        return _ownedTokens[owner];
    }



 **/

contract Ticket721 is ERC721Enumerable, ERC721Mintable {
   using SafeMath for uint256;
   using Counters for Counters.Counter;

    address _factory_address;

    //events
    event TicketBought(address indexed visitor_wallet,uint256 indexed event_id, uint256 indexed ticket_id);
    event TicketBoughtHuman(address visitor_wallet,uint256 event_id, uint256 ticket_id);
    event TicketFulfilled(address indexed visitor_wallet,uint256 indexed event_id, uint256 indexed ticket_id);   // FIXME: add date to event?
    event TicketFulfilledHuman(address visitor_wallet,uint256 event_id, uint256 ticket_id);
    event EventIdReserved(address indexed ticket_sale, uint256 indexed event_id);
    event EventIdReservedHuman(address ticket_sale, uint256 event_id);

    // Global counters for ticket_id and event_id
    Counters.Counter _ticket_id_count;
    //Counters.Counter _ticket_type;
    Counters.Counter _event_id_count;

    // Ticket lifecycle
    enum TicketState {Paid, Fulfilled, Cancelled}

    // map from event id to ticketsale address
    // TIP: ticket type = array.length
    mapping(uint256 => address[]) public eventsales;         // FIXME: one event could have a few ticketsale (for a different ticket types), so change it to address[] (?)
    // map from event id to ticket ids
    mapping (uint256 => uint256[]) ticketIds;
    // map fron token ID to its index in ticketIds
    mapping (uint256 => uint256) ticketIndex;
    // map from ticket id to ticket info
    mapping (uint256 => TicketInfo) ticketInfoStorage;
    // map from sale address to organizer
    mapping(address => address) retailers;


    /**
   * Ticket information
   */

    struct TicketInfo {
    //string description;
    //uint price;
    TicketState state;
   // Counters.Counter ticket_type;
    uint ticket_type;
    //string fulfilmentURI;
    //address retailer;
  }

   // TicketInfo[] internal ticketStorage;



    //FIXME: invoke constructor from 721(?)
    constructor(address factory_address) public {
       // _factory_address = factory_address;
       // addMinter(msg.sender);
    }

    // FIXME: approve for ticketsale, not factory
    function setApprovalForEvent(address _owner, address ticketsale) internal{
        bool approved;
        _operatorApprovals[_owner][ticketsale] = approved;
        emit ApprovalForAll(_owner, ticketsale, approved);
    }

    function _transferFromTicket(address from, address to, uint256 tokenId) public {
        super._transferFrom(from, to, tokenId);
    }

    // TODO - check for event_id already existed
    function reserveEventId(address orginizer) public returns(uint256 event_id){
        _event_id_count.increment();
        event_id = _event_id_count.current();
      //  eventsales[event_id] = msg.sender;
        eventsales[event_id].push(msg.sender);
        retailers[msg.sender] = orginizer;
       // addMinter(msg.sender);
        emit EventIdReserved(msg.sender,event_id);
        emit EventIdReservedHuman(msg.sender,event_id);
        return event_id;
    }

    // plug additional sale for selling different types of ticket by one event
    function plugSale(uint256 event_id, address orginizer) public returns(uint) {
        address[] memory _sales = eventsales[event_id];
        address _sale = _sales[0];
        require(retailers[_sale] == orginizer, "only orginizer can plug sale");
        eventsales[event_id].push(msg.sender);
        uint type_count = getTicketTypeCount(event_id);
        return type_count;
    }

    //TODO - return ticketIDs(?)
    // FIXME: Revert reason: MinterRole: caller does not have the Minter role
    function buyTicket(address buyer, uint256 ticketAmount, uint256 event_id, uint _ticket_type) public{
        address[] memory _sales = eventsales[event_id];
        address _sale = _sales[_ticket_type];
        require(_sale == msg.sender, "you should call buyTicket from ticketsale contract");
        for (uint256 i = 0; i < ticketAmount; i++ ){
            _ticket_id_count.increment();
            uint256 ticket_id = _ticket_id_count.current();

            addMinter(msg.sender);
            _mint(buyer,ticket_id);
            ticketInfoStorage[ticket_id] = TicketInfo(TicketState.Paid,_ticket_type);
            ticketIndex[ticket_id] = ticketIds[event_id].length;
            ticketIds[event_id].push(ticket_id);
            // approve for ticketsale (msg.sender = ticketsale)
            approve(msg.sender, ticket_id);
            emit TicketBought(buyer,event_id,ticket_id);
            emit TicketBoughtHuman(buyer,event_id,ticket_id);
        }

    }

    // This function is burning tokens.
    // @deprecated
    /*
    function redeemTicket(address owner,uint256 tokenId, uint256 event_id) public{
        require(eventsales[event_id] == msg.sender, "caller doesn't match with event_id");
        super._burn(owner,tokenId);

       // To prevent a gap in the tokens array, we store the last token in the index of the token to delete, and
       // then delete the last slot (swap and pop).


        uint256 ticket_index = ticketIndex[tokenId];
        uint256 lastTicketIndex = ticketIds[event_id].length.sub(1);

      //  uint256[] storage ticketIdArray = ticketIds[event_id];
      //  uint256 lastTicketId = ticketIdArray[lastTicketIndex];

        uint256 lastTicketId = ticketIds[event_id][lastTicketIndex];


        ticketIds[event_id][ticket_index] = lastTicketId; // Move the last token to the slot of the to-delete token
        ticketIndex[lastTicketId] = ticket_index;         // Update the moved token's index

        ticketIds[event_id].length--;  // remove last element in array
        ticketIndex[tokenId] = 0;


    }
    */


     function redeemTicket(address visitor, uint256 tokenId, uint256 event_id) public{
        address[] memory _sales = eventsales[event_id];
        TicketInfo memory info = ticketInfoStorage[tokenId];
        address _sale = _sales[info.ticket_type];
        require(_sale == msg.sender, "you should call scan from ticketsale contract");
        require(ticketInfoStorage[tokenId].state == TicketState.Paid, "Ticket state must be Paid");
        info.state = TicketState.Fulfilled;
        ticketInfoStorage[tokenId] = info;
        emit TicketFulfilled(visitor,event_id,tokenId);
        emit TicketFulfilledHuman(visitor, event_id,tokenId);
    }


    function getTicketTypeCount(uint256 event_id) public view returns(uint) {
        address[] memory _sales = eventsales[event_id];
        uint ticket_type = _sales.length;
        return ticket_type;
    }

    function getTicketByOwner(address _owner) public view returns(uint256[] memory) {
        uint256[] storage tickets = _tokensOfOwner(_owner);
        return tickets;
    }

    function getTicketSales(uint256 event_id) public view returns(address[] memory) {
        address[] memory sales = eventsales[event_id];
        return sales;
    }

}