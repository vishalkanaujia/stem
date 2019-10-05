#include <iostream>
#include <list>
#include <unordered_map>

using namespace std;

size_t capacity;

class LRUCache
{
public:
    bool Lookup(int key, int *price)
    {
        auto it = hash_table.find(key);
        if (it == hash_table.end())
        {
            return false;
        }

        *price = it->second.second;

        // update the key in the queue
        MoveToFront(key, it);
        return true;
    }

    void Insert(int key, int price)
    {
        auto it = hash_table.find(key);
        if (it != hash_table.end())
        {
            MoveToFront(key, it);
            return;
        }

        if (hash_table.size() == capacity)
            CleanTable();

        lru_queue.emplace_front(key);
        hash_table[key] = {lru_queue.begin(), price};
    }

    bool Erase(int key)
    {
        auto it = hash_table.find(key);
        if (it != hash_table.end())
        {
            lru_queue.erase(it->second.first);
            hash_table.erase(it);
            return true;
        }
        return false;
    }

    void PrintQueue()
    {
        for (auto it : lru_queue)
        {
            cout << it << " ";
        }
        cout << endl;
    }

private:
    typedef unordered_map<int, pair<list<int>::iterator, int>> Table;
    Table hash_table;

    void MoveToFront(int key, const Table::iterator &it)
    {
        // delete the node.
        lru_queue.erase(it->second.first);

        // create a new node at the front.
        lru_queue.emplace_front(key);
        it->second.first = lru_queue.begin();
    }

    void CleanTable()
    {
        hash_table.erase(lru_queue.back());
        lru_queue.pop_back();
    }

    list<int> lru_queue;
};

int main()
{
    capacity = 3;
    LRUCache cache;

    cache.Insert(1, 100);
    cache.Insert(2, 200);
    cache.Insert(3, 300);

    cache.PrintQueue();

    cache.Insert(4, 400);
    cache.PrintQueue();
}