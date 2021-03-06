package main

import (
	"context"
	"os"
	"strconv"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/util/runtime"

	podcontroller "github.com/azopat/pod-controller/pkg/controller"
)

func main() {

	ctx, _ := context.WithCancel(context.Background())
	loggerConfig := zap.NewProductionConfig()

	// general logger
	logger, err := loggerConfig.Build()
	runtime.Must(err)

	podNamespace := "icap-adaptation"

	podCountStr := os.Getenv("POD_COUNT")
	podCount, err := strconv.Atoi(podCountStr)
	if err != nil {
		podCount = 10 // default value
	}

	rs := &podcontroller.RebuildSettings{
		PodCount: podCount,
	}

	ctrl, err := podcontroller.NewPodController(logger, podNamespace, rs)
	if err != nil {
		logger.Panic("Failed to initialise the controller", zap.Error(err))
	}

	ctrl.Run(ctx)

	<-ctx.Done()
}
