package iconst

const (
	KWorkingSiteIDMypoint    = "8854"
	KWorkingSiteIDEpoint     = "20984"
	KWorkingSiteIDMypointInt = 8854
	KWorkingSiteIDEpointInt  = 20984
	KWorkingSiteIDMypointStr = "MYPOINT"
	KWorkingSiteIDEpointStr  = "EPOINT"
)

var MapWorkingSiteID = map[string]struct{}{
	KWorkingSiteIDMypoint: {},
	KWorkingSiteIDEpoint:  {},
}

var MapWorkingSiteIDWorkspace = map[string]int{
	KWorkingSiteIDMypointStr: KWorkingSiteIDMypointInt,
	KWorkingSiteIDEpointStr:  KWorkingSiteIDEpointInt,
}
