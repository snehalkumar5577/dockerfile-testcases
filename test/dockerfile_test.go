package test

import (
	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDockerfile(t *testing.T) {

	javaVersion := "1.8"
	tag := "1.0.0"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	docker.Build(t, "../", buildOptions)

	goOpts := &docker.RunOptions{Command: []string{"java", "-version"}}
	goOutput := docker.Run(t, tag, goOpts)
	assert.Containsf(t, goOutput, javaVersion, javaVersion)
}
