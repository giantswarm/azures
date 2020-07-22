package annotation

// Docs is the docs annotation put into all CRs to link to its CR specific
// documentation. This aims to help understanding all the moving parts within
// the system and how they relate to each other.
const Docs = "giantswarm.io/docs"

// ReleaseNotes defines where to find release notes about the CR at hand.
// The value is expected to be a URI, e. g.
// "https://github.com/giantswarm/releases/tree/master/aws/v11.5.0".
const ReleaseNotes = "giantswarm.io/release-notes"
