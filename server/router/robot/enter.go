package robot

type RouterGroup struct {
	RobUserRouter
	RobAdminLogRouter
	RobCronLogRouter
	RobExchangeRouter
	RobExchangeSymboRouter
	RobPairRouter
}
