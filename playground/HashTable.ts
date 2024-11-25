class HashTable<K, V> {
    private table: { [key: number]: { key: K, value: V }[] } = {};
    private size: number = 0;

    constructor(private initialCapacity: number = 32) {
        this.table = new Array(initialCapacity).fill(null).map(() => []);
    }

    private hash(key: K): number {
        let hashCode = 0;
        if (typeof key === 'string') {
            for (let i = 0; i < key.length; i++) {
                hashCode += key.charCodeAt(i);
            }
        } else {
            hashCode = key as any;
        }
        return hashCode % this.initialCapacity;
    }

    set(key: K, value: V): void {
        const index = this.hash(key);
        const bucket = this.table[index];

        for (let i = 0; i < bucket.length; i++) {
            if (bucket[i].key === key) {
                bucket[i].value = value;
                return;
            }
        }

        bucket.push({ key, value });
        this.size++;
    }

    get(key: K): V | undefined {
        const index = this.hash(key);
        const bucket = this.table[index];

        for (let i = 0; i < bucket.length; i++) {
            if (bucket[i].key === key) {
                return bucket[i].value;
            }
        }

        return undefined;
    }

    remove(key: K): boolean {
        const index = this.hash(key);
        const bucket = this.table[index];

        for (let i = 0; i < bucket.length; i++) {
            if (bucket[i].key === key) {
                bucket.splice(i, 1);
                this.size--;
                return true;
            }
        }

        return false;
    }

    getSize(): number {
        return this.size;
    }
}

const hashTable = new HashTable()
hashTable.set("key1", "value1")
hashTable.set("key2", "value2")
console.log(hashTable.getSize())
console.log(hashTable.get("key2"))