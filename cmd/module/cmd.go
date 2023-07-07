package main

import (
	"context"
	"github.com/shawnbmccarthy/viam-yahboom-transbot-ros/sensors"
	"go.viam.com/rdk/components/sensor"

	"github.com/edaniels/golog"
	"github.com/shawnbmccarthy/viam-yahboom-transbot-ros/imu"
	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/module"
)

func main() {
	err := realMain()
	if err != nil {
		panic(err)
	}
}

func realMain() error {
	ctx := context.Background()
	logger := golog.NewDevelopmentLogger("client")

	myMod, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}

	/*
	 * add modules here (TBD)
	 */
	err = myMod.AddModelFromRegistry(ctx, movementsensor.API, imu.Model)
	err = myMod.AddModelFromRegistry(ctx, sensor.API, sensors.BatteryModel)

	err = myMod.Start(ctx)
	defer myMod.Close(ctx)
	if err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}
