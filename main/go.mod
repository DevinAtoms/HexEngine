module main

go 1.17

require (
	github.com/DevinAtoms/HexEngine/Engine v0.0.0-20220228234711-ea97cfdd5cfa
	github.com/DevinAtoms/HexEngine/HexMath v0.0.0-20220228224128-0e0e6431a868
	github.com/gen2brain/raylib-go/raylib v0.0.0-20220116181443-e4777d30ee99
)

replace (
	github.com/DevinAtoms/HexEngine/Engine v0.0.0-20220228234711-ea97cfdd5cfa => ../Engine
	github.com/DevinAtoms/HexEngine/HexMath v0.0.0-20220228224128-0e0e6431a868 => ../HexMath
)

