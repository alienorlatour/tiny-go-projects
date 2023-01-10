/*
Package money exposes an API to convert an amount from a currency source to a target one.

This is the responsibility from the caller to instantiate a repository with currencies change rates you want to use from repository package.
The converted can be called on a predefined list of currencies with hardcoded precisions for each money.

For float precisions issues, handled amount that can be converted are restricted to a defined range.
*/
package money
