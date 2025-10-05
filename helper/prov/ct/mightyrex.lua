local path = arg[0]:match("(.*[/\\])")
dofile(path.."../../lib/reelgen.lua")

local symset = {
	2, --  1 wild         15000
	2, --  2 scatter      (on 3, 4, 5 reels)
	2, --  3 einiosaurus  1250
	2, --  4 kentrosaurus 1250
	3, --  5 troodon      500
	3, --  6 spinosaurus  500
	3, --  7 cretoxyrhina 300
	3, --  8 ammonite     300
	4, --  9 ace          200
	4, -- 10 king         200
	5, -- 11 queen        150
	5, -- 12 jack         150
}

local neighbours = {
	--1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,
	{ 2, 2, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, }, --  1 wild
	{ 2, 2, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, }, --  2 scatter (on 3, 4, 5 reels)
	{ 1, 1, 2, 1, 1, 1, 0, 0, 0, 0, 0, 0, }, --  3 einiosaurus
	{ 1, 1, 1, 2, 1, 1, 0, 0, 0, 0, 0, 0, }, --  4 kentrosaurus
	{ 1, 1, 1, 1, 2, 1, 0, 0, 0, 0, 0, 0, }, --  5 troodon
	{ 1, 1, 1, 1, 1, 2, 0, 0, 0, 0, 0, 0, }, --  6 spinosaurus
	{ 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, }, --  7 cretoxyrhina
	{ 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, }, --  8 ammonite
	{ 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, }, --  9 ace
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, }, -- 10 king
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, }, -- 11 queen
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, }, -- 12 jack
}

math.randomseed(os.time())
printreel(makereel(symset, neighbours))

local p, n = 1/2133.3, 15
print(p*n)
print(1 - (1-p)^n)
