-- zero particles Pause Linear Tween
-- self.linearTween = require("zeroParticles/helpers/zpPLinearTween").new( msg.url() );

local M = {}

function M.new(sender)
	local state = {
		sender = sender, 
		tweenValues = {},
		time = 0,
		state = 1, -- 1: fadein, 2: pause, 3: fadeout
		value = 0,
		vec = vmath.vector3()
	}
	
	return setmetatable(state, { __index = M })
end

function M:setValues(startValue, tweenInDuration, tweenInValue, pauseDuration, tweenOutDuration, tweenOutValue)
	self.tweenValues.startValue = startValue;
	self.tweenValues.tweenInDuration = tweenInDuration;
	self.tweenValues.tweenInValue = tweenInValue;
	self.tweenValues.pauseDuration = pauseDuration;
	self.tweenValues.tweenOutDuration = tweenOutDuration;
	self.tweenValues.tweenOutValue = tweenOutValue;

	self.tweenValues.afterPauseTime = self.tweenValues.tweenInDuration + self.tweenValues.pauseDuration;
	self.tweenValues.completeDuration = self.tweenValues.tweenInDuration + self.tweenValues.pauseDuration + self.tweenValues.tweenOutDuration; 
	
	self.value = self.tweenValues.startValue;
end

function M:update(dt)
	self.time = self.time + dt;
	if(self.state == 1)then
		self.value = self.tweenValues.startValue + (self.tweenValues.tweenInValue - self.tweenValues.startValue) * self.time / self.tweenValues.tweenInDuration;
		if(self.time >= self.tweenValues.tweenInDuration)then
			self.value = self.tweenValues.tweenInValue;
			self.state = 2;
		end
	end

	if(self.state == 2)then
		if(self.time >= self.tweenValues.afterPauseTime)then
			self.state = 3;
		end
	end

	if(self.state == 3)then
		self.value = self.tweenValues.tweenInValue + (self.tweenValues.tweenOutValue - self.tweenValues.tweenInValue) * (self.time - self.tweenValues.afterPauseTime) / self.tweenValues.tweenOutDuration;
		if(self.time >= self.tweenValues.completeDuration)then
			self.value = self.tweenValues.tweenOutValue;
			self.state = 4;
		end
	end
end

function M:calcValueByTime(time)
	if(time < self.tweenValues.tweenInDuration)then
		return self.tweenValues.startValue + time * (self.tweenValues.tweenInValue - self.tweenValues.startValue) / self.tweenValues.tweenInDuration;
		
	elseif(time < self.tweenValues.afterPauseTime)then
		return self.tweenValues.tweenInValue;
		
	elseif(time < self.tweenValues.completeDuration)then
		self.value = self.tweenValues.tweenInValue + (self.tweenValues.tweenOutValue - self.tweenValues.tweenInValue) * (time - self.tweenValues.afterPauseTime) / self.tweenValues.tweenOutDuration;
		
	else
		return self.tweenValues.tweenOutValue;
	end
end

function M:currentValue()
	return self.value;
end

return M