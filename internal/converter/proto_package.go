package converter

import (
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// ProtoPackage describes a package of Protobuf, which is an container of message types.
type ProtoPackage struct {
	name     string
	parent   *ProtoPackage
	children map[string]*ProtoPackage
	types    map[string]*descriptor.DescriptorProto
	enums    map[string]*descriptor.EnumDescriptorProto
}

func (c *Converter) lookupType(pkg *ProtoPackage, name string) (*descriptor.DescriptorProto, string, bool) {
	if strings.HasPrefix(name, ".") {
		return c.relativelyLookupType(globalPkg, name[1:])
	}

	for ; pkg != nil; pkg = pkg.parent {
		if desc, pkgName, ok := c.relativelyLookupType(pkg, name); ok {
			return desc, pkgName, ok
		}
	}
	return nil, "", false
}

func (c *Converter) lookupEnum(pkg *ProtoPackage, name string) (*descriptor.EnumDescriptorProto, string, bool) {
	if strings.HasPrefix(name, ".") {
		return c.relativelyLookupEnum(globalPkg, name[1:])
	}

	for ; pkg != nil; pkg = pkg.parent {
		if desc, pkgName, ok := c.relativelyLookupEnum(pkg, name); ok {
			return desc, pkgName, ok
		}
	}
	return nil, "", false
}

func (c *Converter) relativelyLookupType(pkg *ProtoPackage, name string) (*descriptor.DescriptorProto, string, bool) {
	components := strings.SplitN(name, ".", 2)
	switch len(components) {
	case 0:
		c.logger.Debug("empty message name")
		return nil, "", false
	case 1:
		found, ok := pkg.types[components[0]]
		return found, pkg.name, ok
	case 2:
		c.logger.Tracef("Looking for %s in %s at %s (%v)", components[1], components[0], pkg.name, pkg)
		if child, ok := pkg.children[components[0]]; ok {
			found, pkgName, ok := c.relativelyLookupType(child, components[1])
			return found, pkgName, ok
		}
		if msg, ok := pkg.types[components[0]]; ok {
			found, ok := c.relativelyLookupNestedType(msg, components[1])
			return found, pkg.name, ok
		}
		c.logger.WithField("component", components[0]).WithField("package_name", pkg.name).Info("No such package nor message in package")
		return nil, "", false
	default:
		c.logger.Error("Failed to lookup type")
		return nil, "", false
	}
}

func (c *Converter) relativelyLookupEnum(pkg *ProtoPackage, name string) (*descriptor.EnumDescriptorProto, string, bool) {
	components := strings.SplitN(name, ".", 2)
	switch len(components) {
	case 0:
		c.logger.Debug("empty enum name")
		return nil, "", false
	case 1:
		found, ok := pkg.enums[components[0]]
		return found, pkg.name, ok
	case 2:
		c.logger.Tracef("Looking for %s in %s at %s (%v)", components[1], components[0], pkg.name, pkg)
		if child, ok := pkg.children[components[0]]; ok {
			found, pkgName, ok := c.relativelyLookupEnum(child, components[1])
			return found, pkgName, ok
		}
		// if msg, ok := pkg.types[components[0]]; ok {
		// 	found, ok := c.relativelyLookupNestedEnum(msg, components[1])
		// 	return found, pkg.name, ok
		// }
		c.logger.WithField("component", components[0]).WithField("package_name", pkg.name).Info("No such package nor message in package")
		return nil, "", false
	default:
		c.logger.Error("Failed to lookup type")
		return nil, "", false
	}
}

func (c *Converter) relativelyLookupPackage(pkg *ProtoPackage, name string) (*ProtoPackage, bool) {
	components := strings.Split(name, ".")
	for _, c := range components {
		var ok bool
		pkg, ok = pkg.children[c]
		if !ok {
			return nil, false
		}
	}
	return pkg, true
}
