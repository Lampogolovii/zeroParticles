-- self.zpConsts = require("zeroParticles/zpConsts")

local M = {}

M.MSG_EMIT = hash("zpEmit")

M.MSG_CALCPOSITION = hash( "sendCalculatePosition" ) -- ask shape to send position back to emitter
M.MSG_POSITIONCALCULATED = hash( "positionCalculated" ); -- respone from the shape
M.MSG_ADDPARTICLE = hash("addParticle")
M.MSG_PARTICLEDELETED = hash("particleDeleted")
M.MSG_REMOVEPARTICLES = hash("removeParticles")

M.MSG_PARTICLE_INITED = hash("particleInited")
M.MSG_PARTICLE_BUILDERDATA = hash("particleBuilderData")

-- Particle options:

-- emitterPosition (zpPhysicsBorderHorizontal)
-- value (zpLabelTextValue, zpRectExpanding)
-- image (zpImageName)
-- builderURL (zpDeleteDuration)
-- target (zpTrailSin)
-- duration (zpTrailSin)
-- lifeDuration (zpDeleteDuration)
-- trimPivots (zpTrimFlipbook)
-- custom (use your own data)
-- tint (zpLabelTextValue, zpRectExpanding)
-- scale (zpLabelTextValue)

M.MSG_PARTICLE_OPTIONS = hash("particleOptions")


-- fillTable method
M.fillTable = function (options, emitterData)
	local data = options;
	if(data == nil)then
		data = {}
	end
	for key, value in pairs(emitterData) do
		if(data[key] == nil)then
			data[key] = value;
		end
	end
	return data;
end

return M