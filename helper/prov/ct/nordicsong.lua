local path = arg[0]:match("(.*[/\\])")
dofile(path.."../../lib/reelgen.lua")

local symset1 = {
	1, --  1 wild    (2, 3, 4, 5 reels only)
	2, --  2 scatter (1, 3, 5 reels only)
	1, --  3 man     1000
	4, --  4 woman   500
	3, --  5 owl     400
	3, --  6 dog     400
	4, --  7 ace     200
	4, --  8 king    200
	4, --  9 queen   100
	4, -- 10 jack    100
	4, -- 11 ten     100
}

local neighbours = {
	--1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,
	{ 2, 2, 1, 1, 1, 1, 0, 0, 0, 0, 0,}, --  1 wild    (2, 3, 4, 5 reels only)
	{ 2, 2, 1, 1, 1, 1, 0, 0, 0, 0, 0,}, --  2 scatter (1, 3, 5 reels only)
	{ 1, 1, 2, 1, 1, 1, 0, 0, 0, 0, 0,}, --  3 man
	{ 1, 1, 1, 2, 1, 1, 0, 0, 0, 0, 0,}, --  4 woman
	{ 1, 1, 1, 1, 2, 1, 0, 0, 0, 0, 0,}, --  5 owl
	{ 1, 1, 1, 1, 1, 2, 0, 0, 0, 0, 0,}, --  6 dog
	{ 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0,}, --  7 ace
	{ 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0,}, --  8 king
	{ 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0,}, --  9 queen
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0,}, -- 10 jack
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2,}, -- 11 ten
}

local symset2 = {
	0, --  1 wild    (2, 3, 4, 5 reels only)
	0, --  2 scatter (1, 3, 5 reels only)
	4, --  3 man
	0, --  4 woman
	0, --  5 owl
	0, --  6 dog
	0, --  7 ace
	0, --  8 king
	0, --  9 queen
	0, -- 10 jack
	0, -- 11 ten
}

local chunklen = {
	4, --  1 wild    (2, 3, 4, 5 reels only)
	1, --  2 scatter (1, 3, 5 reels only)
	4, --  3 man
	4, --  4 woman
	4, --  5 owl
	4, --  6 dog
	4, --  7 ace
	4, --  8 king
	4, --  9 queen
	4, -- 10 jack
	4, -- 11 ten
}

math.randomseed(os.time())

local function batchreel(comment)
	local reel1, iter1 = makereel(symset1, neighbours)
	local reel2, iter2 = makereelhot(symset2, 3, {[2]=true}, chunklen, false)
	print(comment)
	if iter1 >= 1000 then
		print(string.format("iterations: %d, %d", iter1, iter2))
	end
	printreel(tableglue(reel1, reel2))
end

do
	local n11, n21 = symset1[1], symset2[1]
	symset1[1], symset2[1] = 0, 0
	batchreel "reel 1"
	symset1[1], symset2[1] = n11, n21
end

do
	local n12, n22 = symset1[2], symset2[2]
	symset1[2], symset2[2] = 0, 0
	batchreel "reel 2, 4"
	symset1[2], symset2[2] = n12, n22
end

do
	batchreel "reel 3, 5"
end
