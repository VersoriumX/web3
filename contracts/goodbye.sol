pragma solidity ^0.5.0;

interface PriceSource {
	function latestRoundData() external view returns (uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound);
	function decimals() external view returns (uint8);
// source: https://polygonscan.com/address/0xab594600376ec9fd91f8e885dadf0ce036862de0#code
}
