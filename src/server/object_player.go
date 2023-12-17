package server

import "unsafe"

type PlayerUpdateData struct {
	Field0           uint32             // 0, 0
	ManaCur          uint16             // 1, 4
	ManaPrev         uint16             // 1, 6
	ManaMax          uint16             // 2, 8
	Field2_1         uint16             // 2, 10
	Field3           uint32             // 3, 12 // TODO: [32]uint16 ?
	Field4           uint32             // 4, 16
	Field5           uint32             // 5, 20
	Field6           uint32             // 6, 24
	Field7           uint32             // 7, 28
	Field8           uint32             // 8, 32
	Field9           uint32             // 9, 36
	Field10          uint32             // 10, 40
	Field11          uint32             // 11, 44
	Field12          uint32             // 12, 48
	Field13          uint32             // 13, 52
	Field14          uint32             // 14, 56
	Field15          uint32             // 15, 60
	Field16          uint32             // 16, 64
	Field17          uint32             // 17, 68
	Field18          uint32             // 18, 72
	Field19_0        uint16             // 19, 76
	Field19_1        uint16             // 19, 78
	Field20_0        uint16             // 20, 80
	Field20_1        uint16             // 20, 82
	Field21          float32            // 21, 84
	State            uint8              // 22, 88
	State2           uint8              // 22, 89
	Field22_2        uint8              // 22, 90
	Field22_3        uint8              // 22, 91
	Field23          uint32             // 23, 92
	Field24          uint32             // 24, 96
	Field25          uint32             // 25, 100
	Field26          *Object            // 26, 104
	Field27          *Object            // 27, 108, TODO: weapon?
	Field28          uint32             // 28, 112
	Field29          [4]*Object         // 29, 116, TODO: teleport markers? traps?
	HarpoonTarg      *Object            // 33, 132
	HarpoonBolt      *Object            // 34, 136
	Harpoon35        uint32             // 35, 140
	HarpoonTargX     float32            // 36, 144
	HarpoonTargY     float32            // 37, 148
	HarpoonFrame     uint32             // 38, 152
	Field39          uint32             // 39, 156
	Field40_0        uint16             // 40, 160
	Field40_1        uint16             // 40, 162
	Field41          uint32             // 41, 164
	Field42          uint32             // 42, 168, TODO: [?]*Object ?
	Field43          uint32             // 43, 172
	Field44          uint32             // 44, 176
	Field45          uint32             // 45, 180
	SpellPhonemeLeaf *PhonemeLeaf       // 46, 184
	Field47_0        uint8              // 47, 188
	Field47_1        uint8              // 47, 189
	Field47_2        uint16             // 47, 190
	TrapSpells       [5]uint32          // 48, 192
	TrapSpellsCnt    uint32             // 53, 212
	SpellCastStart   uint32             // 54, 216
	Field55          int                // 55, 220, TODO: spell-related? x coord?
	Field56          int                // 56, 224, TODO: spell-related? y coord?
	Field57          uint32             // 57, 228
	Field58          uint32             // 58, 232
	Field59_0        uint8              // 59, 236, TODO: frame index?
	Field59_1        uint8              // 59, 237
	Field59_2        uint16             // 59, 238
	Field60          uint32             // 60, 240, TODO: movement flags?
	CurTraps         uint32             // 61, 244
	Field62          uint32             // 62, 248
	Field63          uint32             // 63, 252
	Field64          uint32             // 64, 256
	Field65          uint32             // 65, 260
	Field66          uint32             // 66, 264
	Field67          uint32             // 67, 268
	Field68          uint32             // 68, 272
	Player           *Player            // 69, 276
	Trade70          *TradeSession      // 70, 280
	DialogWith       *Object            // 71, 284
	CursorObj        *Object            // 72, 288
	UpdateDataBot    *MonsterUpdateData // 73, 292
	Field74          uint32             // 74, 296
	Field75          uint32             // 75, 300
	Field76          uint32             // 76, 304
	Field77          unsafe.Pointer     // 77, 308
	Field78          uint32             // 78, 312
	Field79          uint32             // 79, 316
	Field80          uint32             // 80, 320
	Field81          uint32             // 81, 324
	Field82          uint32             // 82, 328
	Field83          uint32             // 83, 332
	Field84          uint32             // 84, 336
	Field85          uint32             // 85, 340
	Field86          uint32             // 86, 344
	Field87          uint32             // 87, 348
	Field88          uint32             // 88, 352
	Field89          uint32             // 89, 356
	Field90          uint32             // 90, 360
	Field91          uint32             // 91, 364
	Field92          uint32             // 92, 368
	Field93          uint32             // 93, 372
	Field94          uint32             // 94, 376
	Field95          uint32             // 95, 380
	Field96          uint32             // 96, 384
	Field97          uint32             // 97, 388
	Field98          uint32             // 98, 392
	Field99          uint32             // 99, 396
	Field100         uint32             // 100, 400
	Field101         uint32             // 101, 404
	Field102         uint32             // 102, 408
	Field103         uint32             // 103, 412
	Field104         uint32             // 104, 416
	Field105         uint32             // 105, 420
	Field106         uint32             // 106, 424
	Field107         uint32             // 107, 428
	Field108         uint32             // 108, 432
	Field109         uint32             // 109, 436
	Field110         uint32             // 110, 440
	Field111         uint32             // 111, 444
	Field112         uint32             // 112, 448
	Field113         uint32             // 113, 452
	Field114         uint32             // 114, 456
	Field115         uint32             // 115, 460
	Field116         uint32             // 116, 464
	Field117         uint32             // 117, 468
	Field118         uint32             // 118, 472
	Field119         uint32             // 119, 476
	Field120         uint32             // 120, 480
	Field121         uint32             // 121, 484
	Field122         uint32             // 122, 488
	Field123         uint32             // 123, 492
	Field124         uint32             // 124, 496
	Field125         uint32             // 125, 500
	Field126         uint32             // 126, 504
	Field127         uint32             // 127, 508
	Field128         uint32             // 128, 512
	Field129         uint32             // 129, 516
	Field130         uint32             // 130, 520
	Field131         uint32             // 131, 524
	Field132         uint32             // 132, 528
	Field133         uint32             // 133, 532
	Field134         uint32             // 134, 536
	Field135         uint32             // 135, 540
	Field136         uint32             // 136, 544
	Field137         uint32             // 137, 548, TODO: some timestamp
	Field138         uint32             // 138, 552
}
