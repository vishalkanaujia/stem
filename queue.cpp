#include <iostream>
#include <list>
#include <stack>
#include <algorithm>

using namespace std;

class Queue
{
private:
    list<int> data;

public:
    void Enqueue(int x)
    {
        data.__emplace_back(x);
    }

    int Dequeue()
    {
        if (data.empty())
        {
            throw length_error("empty queue");
        }
        int val = data.front();
        data.pop_front();
        return val;
    }

    int Max() const
    {
        if (data.empty())
        {
            throw length_error("can't find max, empty queue");
        }
        return *max_element(data.begin(), data.end());
    }
};

class QueueWithStacks
{
public:
    void Enqueue(int x)
    {
        enq_.push(x);
    }

    int Dequeue()
    {
        if (deq_.empty())
        {
            // transfer elements from ena_ stack
            while (!enq_.empty())
            {
                deq_.push(enq_.top());
                enq_.pop();
            }
        }

        int val = deq_.top();
        deq_.pop();
        return val;
    }

private:
    stack<int> enq_, deq_;
};

int main()
{
    Queue q;
    q.Enqueue(10);
    q.Enqueue(20);
    cout << q.Dequeue() << endl;
}