# Mongo Filter Structure
![Test workflow](https://github.com/jobsearch-demos/mongo-filter-struct/actions/workflows/test.yml/badge.svg)

This is a simple filter that will scan the golang struct type and
generate a **mongo query** based on the struct `fields`, their `tags` and
their `values`.

## Currently supported logic:

 - Zero Level (no nested structs) scan
 - Nested structs scan
 - JOINs from different collections (using $lookup)
 - Merge operations (merging the fields with the same name) with several logic operators (AND, OR, XOR, NOT)
 - Currently provided operators:
   - $eq
   - $ne
   - $gt
   - $gte
   - $lt
   - $lte
   - $in
   - $nin
   - $regex

## Customization

You can customize all the `policies` (i.e. merge and join policies) and `operators` by implementing the **interfaces**
To add a new operator you need to implement the `IOperator` interface and add it to the `IOperatorMap`
To add a new policy you need to implement the `IPolicy` interface and add it to the `IPolicyMap`
