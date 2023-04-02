# Blaid

A simple, more specific alternative to RAID that prevents unnecessary, redundant storage of non-critical data as well as excessive disk space and IO operations dedicated to parity.

Blaid operates at the directory level to store data redundantly across multiple hardware devices according to user-defined priority.

The user allocates only the directories they wish to store redundantly to "Tier1" or "Tier2" according to their defined criticality. 

Blaid then mirrors the data across two or three hardware devices according to the "Tier" to which the directory is assigned.

Any directories not assigned to a "Tier" are ignored by Blaid and stored according to the operating system configuration.
