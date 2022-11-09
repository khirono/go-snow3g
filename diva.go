package snow3g

var DivA = [...]uint32{
	0x00000000, 0x180f40cd, 0x301e8033, 0x2811c0fe,
	0x603ca966, 0x7833e9ab, 0x50222955, 0x482d6998,
	0xc078fbcc, 0xd877bb01, 0xf0667bff, 0xe8693b32,
	0xa04452aa, 0xb84b1267, 0x905ad299, 0x88559254,
	0x29f05f31, 0x31ff1ffc, 0x19eedf02, 0x01e19fcf,
	0x49ccf657, 0x51c3b69a, 0x79d27664, 0x61dd36a9,
	0xe988a4fd, 0xf187e430, 0xd99624ce, 0xc1996403,
	0x89b40d9b, 0x91bb4d56, 0xb9aa8da8, 0xa1a5cd65,
	0x5249be62, 0x4a46feaf, 0x62573e51, 0x7a587e9c,
	0x32751704, 0x2a7a57c9, 0x026b9737, 0x1a64d7fa,
	0x923145ae, 0x8a3e0563, 0xa22fc59d, 0xba208550,
	0xf20decc8, 0xea02ac05, 0xc2136cfb, 0xda1c2c36,
	0x7bb9e153, 0x63b6a19e, 0x4ba76160, 0x53a821ad,
	0x1b854835, 0x038a08f8, 0x2b9bc806, 0x339488cb,
	0xbbc11a9f, 0xa3ce5a52, 0x8bdf9aac, 0x93d0da61,
	0xdbfdb3f9, 0xc3f2f334, 0xebe333ca, 0xf3ec7307,
	0xa492d5c4, 0xbc9d9509, 0x948c55f7, 0x8c83153a,
	0xc4ae7ca2, 0xdca13c6f, 0xf4b0fc91, 0xecbfbc5c,
	0x64ea2e08, 0x7ce56ec5, 0x54f4ae3b, 0x4cfbeef6,
	0x04d6876e, 0x1cd9c7a3, 0x34c8075d, 0x2cc74790,
	0x8d628af5, 0x956dca38, 0xbd7c0ac6, 0xa5734a0b,
	0xed5e2393, 0xf551635e, 0xdd40a3a0, 0xc54fe36d,
	0x4d1a7139, 0x551531f4, 0x7d04f10a, 0x650bb1c7,
	0x2d26d85f, 0x35299892, 0x1d38586c, 0x053718a1,
	0xf6db6ba6, 0xeed42b6b, 0xc6c5eb95, 0xdecaab58,
	0x96e7c2c0, 0x8ee8820d, 0xa6f942f3, 0xbef6023e,
	0x36a3906a, 0x2eacd0a7, 0x06bd1059, 0x1eb25094,
	0x569f390c, 0x4e9079c1, 0x6681b93f, 0x7e8ef9f2,
	0xdf2b3497, 0xc724745a, 0xef35b4a4, 0xf73af469,
	0xbf179df1, 0xa718dd3c, 0x8f091dc2, 0x97065d0f,
	0x1f53cf5b, 0x075c8f96, 0x2f4d4f68, 0x37420fa5,
	0x7f6f663d, 0x676026f0, 0x4f71e60e, 0x577ea6c3,
	0xe18d0321, 0xf98243ec, 0xd1938312, 0xc99cc3df,
	0x81b1aa47, 0x99beea8a, 0xb1af2a74, 0xa9a06ab9,
	0x21f5f8ed, 0x39fab820, 0x11eb78de, 0x09e43813,
	0x41c9518b, 0x59c61146, 0x71d7d1b8, 0x69d89175,
	0xc87d5c10, 0xd0721cdd, 0xf863dc23, 0xe06c9cee,
	0xa841f576, 0xb04eb5bb, 0x985f7545, 0x80503588,
	0x0805a7dc, 0x100ae711, 0x381b27ef, 0x20146722,
	0x68390eba, 0x70364e77, 0x58278e89, 0x4028ce44,
	0xb3c4bd43, 0xabcbfd8e, 0x83da3d70, 0x9bd57dbd,
	0xd3f81425, 0xcbf754e8, 0xe3e69416, 0xfbe9d4db,
	0x73bc468f, 0x6bb30642, 0x43a2c6bc, 0x5bad8671,
	0x1380efe9, 0x0b8faf24, 0x239e6fda, 0x3b912f17,
	0x9a34e272, 0x823ba2bf, 0xaa2a6241, 0xb225228c,
	0xfa084b14, 0xe2070bd9, 0xca16cb27, 0xd2198bea,
	0x5a4c19be, 0x42435973, 0x6a52998d, 0x725dd940,
	0x3a70b0d8, 0x227ff015, 0x0a6e30eb, 0x12617026,
	0x451fd6e5, 0x5d109628, 0x750156d6, 0x6d0e161b,
	0x25237f83, 0x3d2c3f4e, 0x153dffb0, 0x0d32bf7d,
	0x85672d29, 0x9d686de4, 0xb579ad1a, 0xad76edd7,
	0xe55b844f, 0xfd54c482, 0xd545047c, 0xcd4a44b1,
	0x6cef89d4, 0x74e0c919, 0x5cf109e7, 0x44fe492a,
	0x0cd320b2, 0x14dc607f, 0x3ccda081, 0x24c2e04c,
	0xac977218, 0xb49832d5, 0x9c89f22b, 0x8486b2e6,
	0xccabdb7e, 0xd4a49bb3, 0xfcb55b4d, 0xe4ba1b80,
	0x17566887, 0x0f59284a, 0x2748e8b4, 0x3f47a879,
	0x776ac1e1, 0x6f65812c, 0x477441d2, 0x5f7b011f,
	0xd72e934b, 0xcf21d386, 0xe7301378, 0xff3f53b5,
	0xb7123a2d, 0xaf1d7ae0, 0x870cba1e, 0x9f03fad3,
	0x3ea637b6, 0x26a9777b, 0x0eb8b785, 0x16b7f748,
	0x5e9a9ed0, 0x4695de1d, 0x6e841ee3, 0x768b5e2e,
	0xfedecc7a, 0xe6d18cb7, 0xcec04c49, 0xd6cf0c84,
	0x9ee2651c, 0x86ed25d1, 0xaefce52f, 0xb6f3a5e2,
}
