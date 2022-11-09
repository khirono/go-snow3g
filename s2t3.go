package snow3g

var S2T3 = [...]uint32{
	0x4a6f2525, 0x486c2424, 0xe6957373, 0xcea96767,
	0xc710d7d7, 0x359baeae, 0xb8e45c5c, 0x60503030,
	0x2185a4a4, 0xb55beeee, 0xdcb26e6e, 0xff34cbcb,
	0xfa877d7d, 0x03b6b5b5, 0x6def8282, 0xdf04dbdb,
	0xa145e4e4, 0x75fb8e8e, 0x90d84848, 0x92db4949,
	0x9ed14f4f, 0xbae75d5d, 0xd4be6a6a, 0xf0887878,
	0xe0907070, 0x79f18888, 0xb951e8e8, 0xbee15f5f,
	0xbce25e5e, 0x61e58484, 0xcaaf6565, 0xad4fe2e2,
	0xd901d8d8, 0xbb52e9e9, 0xf13dcccc, 0xb35eeded,
	0x80c04040, 0x5e712f2f, 0x22331111, 0x50782828,
	0xaef95757, 0xcd1fd2d2, 0x319dacac, 0xaf4ce3e3,
	0x94de4a4a, 0x2a3f1515, 0x362d1b1b, 0x1ba2b9b9,
	0x0dbfb2b2, 0x69e98080, 0x63e68585, 0x2583a6a6,
	0x5c722e2e, 0x04060202, 0x8ec94747, 0x527b2929,
	0x0e090707, 0x96dd4b4b, 0x1c120e0e, 0xeb2ac1c1,
	0xa2f35151, 0x3d97aaaa, 0x7bf28989, 0xc115d4d4,
	0xfd37caca, 0x02030101, 0x8cca4646, 0x0fbcb3b3,
	0xb758efef, 0xd30edddd, 0x88cc4444, 0xf68d7b7b,
	0xed2fc2c2, 0xfe817f7f, 0x15abbebe, 0xef2cc3c3,
	0x57c89f9f, 0x40602020, 0x98d44c4c, 0xc8ac6464,
	0x6fec8383, 0x2d8fa2a2, 0xd0b86868, 0x84c64242,
	0x26351313, 0x01b5b4b4, 0x82c34141, 0xf33ecdcd,
	0x1da7baba, 0xe523c6c6, 0x1fa4bbbb, 0xdab76d6d,
	0x9ad74d4d, 0xe2937171, 0x42632121, 0x8175f4f4,
	0x73fe8d8d, 0x09b9b0b0, 0xa346e5e5, 0x4fdc9393,
	0x956bfefe, 0x77f88f8f, 0xa543e6e6, 0xf738cfcf,
	0x86c54343, 0x8acf4545, 0x62533131, 0x44662222,
	0x6e593737, 0x6c5a3636, 0x45d39696, 0x9d67fafa,
	0x11adbcbc, 0x1e110f0f, 0x10180808, 0xa4f65252,
	0x3a271d1d, 0xaaff5555, 0x342e1a1a, 0xe326c5c5,
	0x9cd24e4e, 0x46652323, 0xd2bb6969, 0xf48e7a7a,
	0x4ddf9292, 0x9768ffff, 0xb6ed5b5b, 0xb4ee5a5a,
	0xbf54ebeb, 0x5dc79a9a, 0x38241c1c, 0x3b92a9a9,
	0xcb1ad1d1, 0xfc827e7e, 0x1a170d0d, 0x916dfcfc,
	0xa0f05050, 0x7df78a8a, 0x05b3b6b6, 0xc4a66262,
	0x8376f5f5, 0x141e0a0a, 0x9961f8f8, 0xd10ddcdc,
	0x06050303, 0x78443c3c, 0x18140c0c, 0x724b3939,
	0x8b7af1f1, 0x19a1b8b8, 0x8f7cf3f3, 0x7a473d3d,
	0x8d7ff2f2, 0xc316d5d5, 0x47d09797, 0xccaa6666,
	0x6bea8181, 0x64563232, 0x2989a0a0, 0x00000000,
	0x0c0a0606, 0xf53bcece, 0x8573f6f6, 0xbd57eaea,
	0x07b0b7b7, 0x2e391717, 0x8770f7f7, 0x71fd8c8c,
	0xf28b7979, 0xc513d6d6, 0x2780a7a7, 0x17a8bfbf,
	0x7ff48b8b, 0x7e413f3f, 0x3e211f1f, 0xa6f55353,
	0xc6a56363, 0xea9f7575, 0x6a5f3535, 0x58742c2c,
	0xc0a06060, 0x936efdfd, 0x4e692727, 0xcf1cd3d3,
	0x41d59494, 0x2386a5a5, 0xf8847c7c, 0x2b8aa1a1,
	0x0a0f0505, 0xb0e85858, 0x5a772d2d, 0x13aebdbd,
	0xdb02d9d9, 0xe720c7c7, 0x3798afaf, 0xd6bd6b6b,
	0xa8fc5454, 0x161d0b0b, 0xa949e0e0, 0x70483838,
	0x080c0404, 0xf931c8c8, 0x53ce9d9d, 0xa740e7e7,
	0x283c1414, 0x0bbab1b1, 0x67e08787, 0x51cd9c9c,
	0xd708dfdf, 0xdeb16f6f, 0x9b62f9f9, 0xdd07dada,
	0x547e2a2a, 0xe125c4c4, 0xb2eb5959, 0x2c3a1616,
	0xe89c7474, 0x4bda9191, 0x3f94abab, 0x4c6a2626,
	0xc2a36161, 0xec9a7676, 0x685c3434, 0x567d2b2b,
	0x339eadad, 0x5bc29999, 0x9f64fbfb, 0xe4967272,
	0xb15decec, 0x66553333, 0x24361212, 0xd50bdede,
	0x59c19898, 0x764d3b3b, 0xe929c0c0, 0x5fc49b9b,
	0x7c423e3e, 0x30281818, 0x20301010, 0x744e3a3a,
	0xacfa5656, 0xab4ae1e1, 0xee997777, 0xfb32c9c9,
	0x3c221e1e, 0x55cb9e9e, 0x43d69595, 0x2f8ca3a3,
	0x49d99090, 0x322b1919, 0x3991a8a8, 0xd8b46c6c,
	0x121b0909, 0xc919d0d0, 0x8979f0f0, 0x65e38686,
}
