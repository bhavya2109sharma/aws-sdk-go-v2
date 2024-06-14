// Code generated by smithy-go-codegen DO NOT EDIT.

// Package budgets provides the API client, operations, and parameter types for
// AWS Budgets.
//
// Use the Amazon Web Services Budgets API to plan your service usage, service
// costs, and instance reservations. This API reference provides descriptions,
// syntax, and usage examples for each of the actions and data types for the Amazon
// Web Services Budgets feature.
//
// Budgets provide you with a way to see the following information:
//
//   - How close your plan is to your budgeted amount or to the free tier limits
//
//   - Your usage-to-date, including how much you've used of your Reserved
//     Instances (RIs)
//
//   - Your current estimated charges from Amazon Web Services, and how much your
//     predicted usage will accrue in charges by the end of the month
//
//   - How much of your budget has been used
//
// Amazon Web Services updates your budget status several times a day. Budgets
// track your unblended costs, subscriptions, refunds, and RIs. You can create the
// following types of budgets:
//
//   - Cost budgets - Plan how much you want to spend on a service.
//
//   - Usage budgets - Plan how much you want to use one or more services.
//
//   - RI utilization budgets - Define a utilization threshold, and receive alerts
//     when your RI usage falls below that threshold. This lets you see if your RIs are
//     unused or under-utilized.
//
//   - RI coverage budgets - Define a coverage threshold, and receive alerts when
//     the number of your instance hours that are covered by RIs fall below that
//     threshold. This lets you see how much of your instance usage is covered by a
//     reservation.
//
// # Service Endpoint
//
// The Amazon Web Services Budgets API provides the following endpoint:
//
//   - https://budgets.amazonaws.com
//
// For information about costs that are associated with the Amazon Web Services
// Budgets API, see [Amazon Web Services Cost Management Pricing].
//
// [Amazon Web Services Cost Management Pricing]: https://aws.amazon.com/aws-cost-management/pricing/
package budgets
