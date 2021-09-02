package reflect

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/internal/diagnostics"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Unknownable is an interface for types that can be explicitly set to known or
// unknown.
type Unknownable interface {
	SetUnknown(context.Context, bool) error
	SetValue(context.Context, interface{}) error
	GetUnknown(context.Context) bool
	GetValue(context.Context) interface{}
}

// NewUnknownable creates a zero value of `target` (or the concrete type it's
// referencing, if it's a pointer) and calls its SetUnknown method.
//
// It is meant to be called through Into, not directly.
func NewUnknownable(ctx context.Context, typ attr.Type, val tftypes.Value, target reflect.Value, opts Options, path *tftypes.AttributePath) (reflect.Value, []*tfprotov6.Diagnostic) {
	var diags []*tfprotov6.Diagnostic
	receiver := pointerSafeZeroValue(ctx, target)
	method := receiver.MethodByName("SetUnknown")
	if !method.IsValid() {
		err := fmt.Errorf("cannot find SetUnknown method on type %s", receiver.Type().String())
		return target, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert value. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
			Attribute: path,
		})
	}
	results := method.Call([]reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(!val.IsKnown()),
	})
	err := results[0].Interface()
	if err != nil {
		var underlyingErr error
		switch e := err.(type) {
		case error:
			underlyingErr = e
		default:
			underlyingErr = fmt.Errorf("unknown error type %T: %v", e, e)
		}
		underlyingErr = fmt.Errorf("reflection error: %w", underlyingErr)
		return target, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert into a value. This is always an error in the provider. Please report the following to the provider developer:\n\n" + underlyingErr.Error(),
			Attribute: path,
		})
	}
	return receiver, diags
}

// FromUnknownable creates an attr.Value from the data in an Unknownable.
//
// It is meant to be called through OutOf, not directly.
func FromUnknownable(ctx context.Context, typ attr.Type, val Unknownable, path *tftypes.AttributePath) (attr.Value, []*tfprotov6.Diagnostic) {
	var diags []*tfprotov6.Diagnostic

	if val.GetUnknown(ctx) {
		tfVal := tftypes.NewValue(typ.TerraformType(ctx), tftypes.UnknownValue)

		if typeWithValidate, ok := typ.(attr.TypeWithValidate); ok {
			diags = append(diags, typeWithValidate.Validate(ctx, tfVal)...)

			if diagnostics.DiagsHasErrors(diags) {
				return nil, diags
			}
		}

		res, err := typ.ValueFromTerraform(ctx, tfVal)
		if err != nil {
			return nil, append(diags, valueFromTerraformErrorDiag(err, path))
		}
		return res, nil
	}
	err := tftypes.ValidateValue(typ.TerraformType(ctx), val.GetValue(ctx))
	if err != nil {
		return nil, append(diags, validateValueErrorDiag(err, path))
	}

	tfVal := tftypes.NewValue(typ.TerraformType(ctx), val.GetValue(ctx))

	if typeWithValidate, ok := typ.(attr.TypeWithValidate); ok {
		diags = append(diags, typeWithValidate.Validate(ctx, tfVal)...)

		if diagnostics.DiagsHasErrors(diags) {
			return nil, diags
		}
	}

	res, err := typ.ValueFromTerraform(ctx, tfVal)
	if err != nil {
		return nil, append(diags, valueFromTerraformErrorDiag(err, path))
	}
	return res, nil
}

// Nullable is an interface for types that can be explicitly set to null.
type Nullable interface {
	SetNull(context.Context, bool) error
	SetValue(context.Context, interface{}) error
	GetNull(context.Context) bool
	GetValue(context.Context) interface{}
}

// NewNullable creates a zero value of `target` (or the concrete type it's
// referencing, if it's a pointer) and calls its SetNull method.
//
// It is meant to be called through Into, not directly.
func NewNullable(ctx context.Context, typ attr.Type, val tftypes.Value, target reflect.Value, opts Options, path *tftypes.AttributePath) (reflect.Value, []*tfprotov6.Diagnostic) {
	var diags []*tfprotov6.Diagnostic
	receiver := pointerSafeZeroValue(ctx, target)
	method := receiver.MethodByName("SetNull")
	if !method.IsValid() {
		err := fmt.Errorf("cannot find SetNull method on type %s", receiver.Type().String())
		return target, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert value. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
			Attribute: path,
		})
	}
	results := method.Call([]reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(val.IsNull()),
	})
	err := results[0].Interface()
	if err != nil {
		var underlyingErr error
		switch e := err.(type) {
		case error:
			underlyingErr = e
		default:
			underlyingErr = fmt.Errorf("unknown error type: %T", e)
		}
		underlyingErr = fmt.Errorf("reflection error: %w", underlyingErr)
		return target, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert into a value. This is always an error in the provider. Please report the following to the provider developer:\n\n" + underlyingErr.Error(),
			Attribute: path,
		})
	}
	return receiver, diags
}

// FromNullable creates an attr.Value from the data in a Nullable.
//
// It is meant to be called through OutOf, not directly.
func FromNullable(ctx context.Context, typ attr.Type, val Nullable, path *tftypes.AttributePath) (attr.Value, []*tfprotov6.Diagnostic) {
	var diags []*tfprotov6.Diagnostic

	if val.GetNull(ctx) {
		tfVal := tftypes.NewValue(typ.TerraformType(ctx), nil)

		if typeWithValidate, ok := typ.(attr.TypeWithValidate); ok {
			diags = append(diags, typeWithValidate.Validate(ctx, tfVal)...)

			if diagnostics.DiagsHasErrors(diags) {
				return nil, diags
			}
		}

		res, err := typ.ValueFromTerraform(ctx, tfVal)
		if err != nil {
			return nil, append(diags, valueFromTerraformErrorDiag(err, path))
		}
		return res, nil
	}
	err := tftypes.ValidateValue(typ.TerraformType(ctx), val.GetValue(ctx))
	if err != nil {
		return nil, append(diags, validateValueErrorDiag(err, path))
	}

	tfVal := tftypes.NewValue(typ.TerraformType(ctx), val.GetValue(ctx))

	if typeWithValidate, ok := typ.(attr.TypeWithValidate); ok {
		diags = append(diags, typeWithValidate.Validate(ctx, tfVal)...)

		if diagnostics.DiagsHasErrors(diags) {
			return nil, diags
		}
	}

	res, err := typ.ValueFromTerraform(ctx, tfVal)
	if err != nil {
		return nil, append(diags, valueFromTerraformErrorDiag(err, path))
	}
	return res, diags
}

// NewValueConverter creates a zero value of `target` (or the concrete type
// it's referencing, if it's a pointer) and calls its FromTerraform5Value
// method.
//
// It is meant to be called through Into, not directly.
func NewValueConverter(ctx context.Context, typ attr.Type, val tftypes.Value, target reflect.Value, opts Options, path *tftypes.AttributePath) (reflect.Value, []*tfprotov6.Diagnostic) {
	var diags []*tfprotov6.Diagnostic
	receiver := pointerSafeZeroValue(ctx, target)
	method := receiver.MethodByName("FromTerraform5Value")
	if !method.IsValid() {
		err := fmt.Errorf("could not find FromTerraform5Type method on type %s", receiver.Type().String())
		return target, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert into a value. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
			Attribute: path,
		})
	}
	results := method.Call([]reflect.Value{reflect.ValueOf(val)})
	err := results[0].Interface()
	if err != nil {
		var underlyingErr error
		switch e := err.(type) {
		case error:
			underlyingErr = e
		default:
			underlyingErr = fmt.Errorf("unknown error type: %T", e)
		}
		underlyingErr = fmt.Errorf("reflection error: %w", underlyingErr)
		return target, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert into a value. This is always an error in the provider. Please report the following to the provider developer:\n\n" + underlyingErr.Error(),
			Attribute: path,
		})
	}
	return receiver, diags
}

// FromValueCreator creates an attr.Value from the data in a
// tftypes.ValueCreator, calling its ToTerraform5Value method and converting
// the result to an attr.Value using `typ`.
//
// It is meant to be called from OutOf, not directly.
func FromValueCreator(ctx context.Context, typ attr.Type, val tftypes.ValueCreator, path *tftypes.AttributePath) (attr.Value, []*tfprotov6.Diagnostic) {
	var diags []*tfprotov6.Diagnostic
	raw, err := val.ToTerraform5Value()
	if err != nil {
		return nil, append(diags, toTerraform5ValueErrorDiag(err, path))
	}
	err = tftypes.ValidateValue(typ.TerraformType(ctx), raw)
	if err != nil {
		return nil, append(diags, validateValueErrorDiag(err, path))
	}
	tfVal := tftypes.NewValue(typ.TerraformType(ctx), raw)

	if typeWithValidate, ok := typ.(attr.TypeWithValidate); ok {
		diags = append(diags, typeWithValidate.Validate(ctx, tfVal)...)

		if diagnostics.DiagsHasErrors(diags) {
			return nil, diags
		}
	}

	res, err := typ.ValueFromTerraform(ctx, tfVal)
	if err != nil {
		return nil, append(diags, valueFromTerraformErrorDiag(err, path))
	}
	return res, diags
}

// NewAttributeValue creates a new reflect.Value by calling the
// ValueFromTerraform method on `typ`. It will return an error if the returned
// `attr.Value` is not the same type as `target`.
//
// It is meant to be called through Into, not directly.
func NewAttributeValue(ctx context.Context, typ attr.Type, val tftypes.Value, target reflect.Value, opts Options, path *tftypes.AttributePath) (reflect.Value, []*tfprotov6.Diagnostic) {
	var diags []*tfprotov6.Diagnostic

	if typeWithValidate, ok := typ.(attr.TypeWithValidate); ok {
		diags = append(diags, typeWithValidate.Validate(ctx, val)...)

		if diagnostics.DiagsHasErrors(diags) {
			return target, diags
		}
	}

	res, err := typ.ValueFromTerraform(ctx, val)
	if err != nil {
		return target, append(diags, valueFromTerraformErrorDiag(err, path))
	}
	if reflect.TypeOf(res) != target.Type() {
		err := fmt.Errorf("Cannot use attr.Value %s, only %s is supported because %T is the type in the schema", target.Type(), reflect.TypeOf(res), typ)
		return target, append(diags, &tfprotov6.Diagnostic{
			Severity:  tfprotov6.DiagnosticSeverityError,
			Summary:   "Value Conversion Error",
			Detail:    "An unexpected error was encountered trying to convert into a Terraform value. This is always an error in the provider. Please report the following to the provider developer:\n\n" + err.Error(),
			Attribute: path,
		})
	}
	return reflect.ValueOf(res), diags
}

// FromAttributeValue creates an attr.Value from an attr.Value. It just returns
// the attr.Value it is passed, but reserves the right in the future to do some
// validation on that attr.Value to make sure it matches the type produced by
// `typ`.
//
// It is meant to be called through OutOf, not directly.
func FromAttributeValue(ctx context.Context, typ attr.Type, val attr.Value, path *tftypes.AttributePath) (attr.Value, []*tfprotov6.Diagnostic) {
	var diags []*tfprotov6.Diagnostic

	if typeWithValidate, ok := typ.(attr.TypeWithValidate); ok {
		tfType := typ.TerraformType(ctx)
		tfVal, err := val.ToTerraformValue(ctx)

		if err != nil {
			return val, append(diags, toTerraformValueErrorDiag(err, path))
		}

		diags = append(diags, typeWithValidate.Validate(ctx, tftypes.NewValue(tfType, tfVal))...)

		if diagnostics.DiagsHasErrors(diags) {
			return val, diags
		}
	}

	return val, diags
}