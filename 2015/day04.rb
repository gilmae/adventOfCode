require 'digest'

def mine starts_with
    index = 0
    hash =""
    while hash[0,starts_with.length] != starts_with
        index +=1
        hash = Digest::MD5.hexdigest("ckczppom#{index}")
    end

    return index
end
puts mine("00000")
puts mine("000000")
