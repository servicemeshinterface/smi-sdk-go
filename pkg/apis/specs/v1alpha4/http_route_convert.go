package v1alpha4

/*
Implementing the hub method is pretty easy -- we just have to add an empty
method called `Hub()` to serve as a
[marker](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/conversion?tab=doc#Hub).
We could also just put this inline in our `cronjob_types.go` file.
*/

// Hub marks this type as a conversion hub.
func (*HTTPRouteGroup) Hub() {}
