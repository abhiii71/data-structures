# Day 2: Hash Functions


## Characteristics of a Good Hash Function:

### Deterministic:

Explanation: The hash function must always produce the same hash value for the same input key.
Reason: This consistency is necessary to ensure that the same key is always mapped to the same bucket.


### Uniform Distribution:

Explanation: Hash values should be evenly distributed across all possible buckets.
Reason: Uniform distribution minimizes collisions, ensuring that no single bucket becomes overloaded with keys.
* 
### Fast Computation:

Explanation: The hash function should be quick to compute.
Reason: Efficient computation of hash values ensures that insertions, deletions, and lookups remain fast, maintaining the overall efficiency of the hash table.

## Common Hash Functions:

### Modular Hashing:

Use Case: Best used for simple applications where keys are uniformly distributed, and the table size is a prime number.
Performance: O(1) for computing the hash value, but performance may degrade with poor choice of table size.

#### Multiplication Method:

Use Case: Suitable for scenarios where the keys are not uniformly distributed or when a good distribution is needed without the need for a prime table size.
Performance: Still O(1), but slightly more computationally intensive due to multiplication and floating-point operations.

#### Universal Hashing:

Use Case: Best used in environments where security and collision resistance are paramount, such as cryptographic applications or where the key distribution is highly unpredictable.
Performance: Slightly slower due to additional random parameter management, but provides superior distribution and collision resistance.

