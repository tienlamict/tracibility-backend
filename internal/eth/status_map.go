package eth

var ProductStatusMap = map[uint8]string{
	0: "Created",
	1: "InProgress",
	2: "Completed",
	3: "Rejected",
}

var StepStatusMap = map[uint8]string{
	0: "Manufactured",
	1: "Inspected",
	2: "Packaged",
	3: "Stored",
	4: "Shipped",
	5: "Received",
	6: "Sold",
	7: "Returned",
	8: "Disposed",
}
