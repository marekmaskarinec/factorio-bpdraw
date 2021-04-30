#!/usr/bin/lua

-- ################################################
-- ### Steals info about entities from factorio ###
-- ################################################

print("###############################################################################")
print("### Make sure that lunajson is installed (`sudo luarocks install lunajson`) ###")
print("### and that you clonned this repo recursively (--recursive).               ###")
print("### Alternatively, run `git submodule update --init --recursive`            ###")
print("### to download submodules after clonning.                                  ###")
print("###############################################################################\n")

luna = require("lunajson")

-- List of files that need to be included for the entities script to be executed
-- When factorio starts, the files are in different locations relative to each other?
-- So I just include all files that entities.lua depends on like this.
-- ==================================================================================
local lualib = {
	"util",
	"circuit-connector-sprites",
	"circuit-connector-generated-definitions"
}

local base = {
	"prototypes.entity.pipecovers",
	"prototypes.entity.transport-belt-pictures",
	"prototypes.entity.assemblerpipes",
	"prototypes.entity.hit-effects",
	"prototypes.entity.sounds",
	"prototypes.entity.movement-triggers",
	"prototypes.entity.spidertron-animations",
	"prototypes.entity.character-animations",
	"prototypes.entity.laser-sounds",
	"prototypes.entity.beacon-animations",
	"prototypes.entity.pump-connector"
}

package.preload["__base__/prototypes/entity/spidertron-light-positions"] = function () return require "factorio-data.base.prototypes.entity.spidertron-light-positions" end

-- load those files with package.preload
-- =====================================
for key,value in ipairs(lualib) do package.preload[value] = function() return require("factorio-data.core.lualib." .. value) end end
for key,value in ipairs(base)   do package.preload[value] = function() return require("factorio-data.base."        .. value) end end

-- fake internal factorio tables
-- =============================
defines = { direction = {} } -- directions in the data tables are all nil but whatever. We don't need those.

function has_value(tab, val) for index, value in ipairs(tab) do if value == val then return true end end return false end

-- fake data object
-- ================
data = { raw = { } }
function data:extend(table)
	for k,v in ipairs(table) do
		if v["flags"] ~= nil and has_value(v["flags"], "player-creation") then
			print("adding " .. v["name"])
			if self.raw[v["type"]] == nil then
				self.raw[v["type"]] = {}
			end
			self.raw[v["type"]][v["name"]] = v
		end
	end
end

-- run the factorio scripts. It writes to the data object
-- ======================================================
require("util")
require("factorio-data.base.prototypes.entity.entities")

-- filter out the raw data table
-- =============================
print()
filtered = {}
for k,v in pairs(data.raw) do
	for key,val in pairs(v) do
		if val["picture"] ~= nil then
			print("filtered add static    " .. val["name"])
			filtered[val["name"]] = {}
			filtered[val["name"]]["picture"] = val["picture"]
		elseif val["animation"] ~= nil then
			print("filtered add animation " .. val["name"])
			filtered[val["name"]] = {}
			filtered[val["name"]]["picture"] = val["animation"]
		elseif val["hand_base_picture"] ~= nil then -- for inserters
			print("filtered add inserter  " .. val["name"])
			filtered[val["name"]] = {}
			filtered[val["name"]]["picture"] = {}
			filtered[val["name"]]["picture"]["rot_layers"] = {}
			filtered[val["name"]]["picture"]["rot_layers"][1] = val["hand_base_picture"]
			filtered[val["name"]]["picture"]["rot_layers"][2] = val["hand_closed_picture"]
			filtered[val["name"]]["picture"]["layers"] = {}
			filtered[val["name"]]["picture"]["layers"][1] = val["platform_picture"]["sheet"]
		end
	end
end

-- encode it into a json
-- =====================
local json = luna.encode(filtered)

-- dump that json into a file
-- ==========================
file = io.open("entities.json", "w+")
io.output(file)
io.write(json)
io.close(file)

print("\n###############")
print  ("### SUCCESS ###")
print  ("###############")
