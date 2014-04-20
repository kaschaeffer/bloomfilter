package main

import (
    "fmt"
    "ioutil" // for reading in a CSV
)

type Payments struct {
    paymentList [][3]int
}

// this needs more explanatory comments...
func (p* Payments) buildInviteMap() {
    
    invitedBy = map[[2]int]int{}

    for i:= 0; i < p.length; i++ {
        paymentId = p[i][0]
        senderId = p[i][1]
        recipientId = p[i][2]

        // 
        if m[recipientId] {
            //DO SOMETHING HERE
            existingPaymentId, exisitingInviterId = m[recipientId]

            if paymentId < existingPaymentId {
                m[recipientId] = paymentId, senderId
            }
        } else {
            m[recipientId] = paymentId, senderId
        }

        if m[senderId] {
            existingPaymentId, exisitingInviterId = m[senderId]

            if paymentId < existingPaymentId {
                m[senderId] = paymentId, senderId
            } 

        } else {
            m[senderId] = paymentId, senderId
        }


    }
}

func (p* Payments) InvitedBy(customerId int) {
    return (&p).invitedBy[customerId]
}

func (p* Payments) InviteChain(customerId int) {
    inviteChain = []int{customerId}

    currentCustomerId = customerId
    for {
        inviter = (&p).InvitedBy(currentCustomerId)
        inviteChain = append(inviteChain, inviter)
        currentCustomerId = inviter
    }
}