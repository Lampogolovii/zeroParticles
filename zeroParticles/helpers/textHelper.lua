-- self.textHelper = require("main/assets/scripts/common/textHelper")
local M = {}

-- int - value that should be converted (only positives)
-- valuesCount - count of symbols after decimal point (
-- examples: 
-- shortInt(1286, 1) -> result will be 1.2k)
-- shortInt(12865, 2) -> result will be 12.86k)
-- shortInt(1286543, 2) -> result will be 1.28m)
-- shortInt(1286543, 3) -> result will be 1.286m)
function M.shortInt(int, valuesCount)
	local suffix = {
		{1, ""},
		{1e3, "k"},
		{1e6, "m"},
		{1e9, "b"},
		{1e12, "t"},
		use = 1
	}
	for i, value in ipairs(suffix) do
		if value[1] <= int then
			suffix.use = i;
		end
	end
	if not valuesCount then valuesCount = 1; end
	valuesCount = math.floor(valuesCount);
	local valuesPow = math.pow(10, valuesCount)
	local resultInt = int / suffix[suffix.use][1];
	local decimals = resultInt - math.floor(resultInt);
	decimals = math.floor(decimals * valuesPow) / valuesPow;
	resultInt = math.floor(resultInt) + decimals;
	return {str = resultInt..""..suffix[suffix.use][2], int = resultInt, suffix = suffix[suffix.use][2]}
end


function M.secondsToTimer(secondsCount)
	local intSeconds = math.ceil(secondsCount);
	local intMinutes = math.floor(intSeconds / 60);
	local result = intMinutes..":";
	intSeconds = intSeconds % 60;
	if(intSeconds < 10)then
		result = result.."0"..intSeconds;
	else
		result = result..intSeconds;
	end
	if(intMinutes < 10)then
		result = "0"..result;
	end
	return result;
end

return M