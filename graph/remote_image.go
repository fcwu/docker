package graph

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/engine"
	"github.com/docker/docker/registry"
)

func (s *TagStore) CmdRemoteImageJson(job *engine.Job) engine.Status {
	if n := len(job.Args); n != 1 {
		return job.Errorf("Usage: %s IMAGE", job.Name)
	}

	var (
		localName   = job.Args[0]
		authConfig  = &registry.AuthConfig{}
		metaHeaders map[string][]string
        imageTag string
	)

	job.GetenvJson("authConfig", authConfig)
	job.GetenvJson("metaHeaders", &metaHeaders)

    if fields := strings.Split(localName, ":"); len(fields) >= 2 {
        localName = fields[0]
        imageTag = fields[1]
    }
	hostname, remoteName, err := registry.ResolveRepositoryName(localName)
	if err != nil {
		return job.Error(err)
	}

	endpoint, err := registry.NewEndpoint(hostname, s.insecureRegistries)
	if err != nil {
		return job.Error(err)
	}

	r, err := registry.NewSession(authConfig, registry.HTTPRequestFactory(metaHeaders), endpoint, true)
	if err != nil {
		return job.Error(err)
	}

	repoData, err := r.GetRepositoryData(remoteName)
	if err != nil {
		if strings.Contains(err.Error(), "HTTP code: 404") {
			return job.Errorf("Error: image %s not found", remoteName)
		}
		// Unexpected HTTP error
		return job.Error(err)
	}

	tagsList, err := r.GetRemoteTags(repoData.Endpoints, remoteName, repoData.Tokens)
	if err != nil {
		log.Errorf("%v", err)
		return job.Error(err)
	}

    var imageId string
	for tag, id := range tagsList {
        if imageTag != "" && imageTag == tag {
            imageId = id
            break
        }
        if imageTag == "" {
            imageId = id
            if tag == "latest" {
                break
            }
        }
	}
    if imageId == "" {
		return job.Errorf("no tag found")
    }

    for _, ep := range repoData.Endpoints {
        var (
            imgJSON []byte
            err     error
        )

        imgJSON, _, err = r.GetRemoteImageJSON(imageId, ep, repoData.Tokens)
        if err != nil {
            return job.Error(err)
        }
        job.Stdout.Write(imgJSON)
        break
    }

	return engine.StatusOK
}
