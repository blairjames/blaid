# Blaid

A simple, more specific alternative to RAID that prevents unnecessary redundant storage of non-critical data as well as excessive disk space and IO operations dedicated to parity.

Blaid operates at the directory level to store data redundantly across multiple hardware devices according to user-defined priority.

The user allocates only the directories they wish to store redundantly to "Tier1" or "Tier2" according to their defined criticality. 

Blaid then mirrors the data across two or three hardware devices according to the "Tier" to which the directory is assigned.

Any directories not assigned to a "Tier" are ignored by Blaid and stored according to the operating system configuration.

#### Blaid runs as a systemd service so can be managed via "systemctl"

#### Logs can be viewed via "journalctl"

<br/>

### Example of operation

#### Hardware Storage Devices:
- Operating System Device (**OS**)
- Redundant Storage Device One (**ONE**)
- Redundant Storage Device Two (**TWO**)

<br/>

#### Example redundancy of specific directories according to user-defined configuration:

#### Example 1 - (Highly critical directory)

**Directory:** "/home/user/Documents/important_documents"

**Configuration:** Assigned to **Tier One**

Hardware devices the directory is mirrored onto:
- OS
- ONE
- TWO

<br/>

#### Example 2 - (Less critical directory)

**Directory:** "/home/user/Pictures/less_critical_photos"

**Configuration:** Assigned to **Tier Two**

Hardware devices the directory is mirrored onto:
- OS
- ONE

<br/>

#### Example 3 - (Directory not configured in Blaid)

**Directory:** "/home/user/Downloads/junk_files"

**Configuration:** Not assigned

Hardware devices the directory is mirrored onto:
- OS (directory is ignored)


