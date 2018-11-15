// Package nagios is a package that is intented to provide a simple, standardized way
// for Nagios checks to return a result.
//
// It does so by defining a Result type, which has an exitCode and a message associated with it.
//
// Implementing this may look like:
// 		package main
//
// 		var (
// 			warning = 80
// 			critical = 90
// 			result = nagios.Result{}
// 		)
//
// 		func main() {
// 			diskUsed := checkDiskUsage()
//
// 			switch {
// 			case diskUsed > critical:
// 				result.CRITICAL
// 			case diskUsed > warning:
// 				result.WARNING
// 			default:
// 				result.OK
// 			}
//
// 			result.Exit()
// 		}
//
//
package nagios
