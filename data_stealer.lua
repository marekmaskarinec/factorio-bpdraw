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

-- list of files that need to be included for the entities script to be executed
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

for key,value in ipairs(lualib) do
	package.preload[value] = function() return require("factorio-data.core.lualib." .. value) end
end

for key,value in ipairs(base) do
	package.preload[value] = function() return require("factorio-data.base." .. value) end
end

function tprint (tbl, indent)
  if not indent then indent = 0 end
  for k, v in pairs(tbl) do
    formatting = string.rep("    ", indent) .. k .. ": "
    if type(v) == "table" then
      print(formatting)
      tprint(v, indent+1)
    else
      print(formatting .. tostring(v))
    end
  end
end

defines = { direction = {} }
data = { raw = { } }

function has_value(tab, val)
    for index, value in ipairs(tab) do
        if value == val then
            return true
        end
    end
    return false
end

function data:extend(table)
	for k,v in ipairs(table) do
		--print("adding " .. v["type"] .. " --- " .. v["name"])
		if v["flags"] ~= nil and has_value(v["flags"], "player-creation") then
			if self.raw[v["type"]] == nil then
				self.raw[v["type"]] = {}
			end
			self.raw[v["type"]][v["name"]] = v
		end
	end
end



-- run the file from factorio. It writes to the data object
require("util")
require("factorio-data.base.prototypes.entity.factorio-logo")
require("factorio-data.base.prototypes.entity.entities")

filtered = {}

for k,v in pairs(data.raw) do
	for key,val in pairs(v) do
		if val["picture"] ~= nil then
			filtered[val["name"]] = {}
			filtered[val["name"]]["picture"] = val["picture"]
		end
	end
end

-- encode it into a json
local json = luna.encode(filtered)

-- dump that json into a file
file = io.open("entities.json", "w+")
io.output(file)
io.write(json)
io.close(file)

print("\n###############")
print("### SUCCESS ###")
print("###############")
